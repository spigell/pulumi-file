package remote

import (
	"crypto/md5" // #nosec
	"fmt"
	"io/fs"
	"strconv"
	"strings"

	"github.com/pkg/sftp"
	p "github.com/pulumi/pulumi-go-provider"
	"golang.org/x/crypto/ssh"
)

func NewSFTPServer(conn *ssh.Client, cmd string, opts ...sftp.ClientOption) (*sftp.Client, error) {
	s, err := conn.NewSession()
	if err != nil {
		return nil, err
	}
	ok, err := s.SendRequest("exec", true, ssh.Marshal(struct{ Command string }{cmd}))
	if err == nil && !ok {
		err = fmt.Errorf("sftp: command %v failed", cmd)
	}
	if err != nil {
		return nil, err
	}
	pw, err := s.StdinPipe()
	if err != nil {
		return nil, err
	}
	pr, err := s.StdoutPipe()
	if err != nil {
		return nil, err
	}

	return sftp.NewClientPipe(pr, pw, opts...)
}

func NewSFTPClient(client *ssh.Client, sudo bool, sftpServerPath string) (*sftp.Client, error) {
	var conn *sftp.Client
	var err error
	if sudo {
		conn, err = NewSFTPServer(client, "sudo -u root -n "+sftpServerPath)
		if err != nil {
			return nil, fmt.Errorf("creating a sudo sftp client: %w", err)
		}
	} else {
		conn, err = sftp.NewClient(client)
		if err != nil {
			return nil, fmt.Errorf("creating a normal sftp client: %w", err)
		}
	}

	return conn, nil
}

func (f *FileOutputs) deploy(ctx p.Context, params FileInputs) error {
	parsedPerm, err := parsePermissions(*params.Permissions)
	if err != nil {
		return fmt.Errorf("parsing permissions: %w", err)
	}

	client, err := f.Connection.Dial(ctx)
	if err != nil {
		return err
	}

	conn, err := NewSFTPClient(client, *params.UseSudo, *params.SFTPPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	dst, err := conn.Create(*params.Path)
	if err != nil {
		return err
	}

	err = dst.Chmod(parsedPerm)
	if err != nil {
		return err
	}

	reader := strings.NewReader(*params.Content)
	_, err = dst.ReadFrom(reader)
	if err != nil {
		return err
	}

	f.BaseOutputs = BaseOutputs{
		// #nosec
		MD5Sum: fmt.Sprintf("%x", md5.Sum([]byte(*params.Content))),
	}
	return nil
}

func (f *FileOutputs) delete(ctx p.Context, params FileOutputs) error {
	client, err := f.Connection.Dial(ctx)
	if err != nil {
		return err
	}

	conn, err := NewSFTPClient(client, *params.UseSudo, *params.SFTPPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.Remove(*params.Path)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileOutputs) read(ctx p.Context, input FileInputs) (*FileOutputs, error) {
	client, err := f.Connection.Dial(ctx)
	if err != nil {
		return nil, err
	}

	conn, err := NewSFTPClient(client, *input.UseSudo, *input.SFTPPath)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	file, err := conn.Open(*input.Path)
	if err != nil {
		return nil, err
	}

	buf := new(strings.Builder)

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	_, err = file.WriteTo(buf)
	if err != nil {
		return nil, err
	}

	return &FileOutputs{
		BaseOutputs: BaseOutputs{
			// #nosec
			MD5Sum: fmt.Sprintf("%x", md5.Sum([]byte(buf.String()))),
		},
		FileInputs: FileInputs{
			Connection:  input.Connection,
			UseSudo:     input.UseSudo,
			SFTPPath:    input.SFTPPath,
			Path:        input.Path,
			Content:     strPtr(buf.String()),
			Permissions: strPtr(strconv.FormatUint(uint64(stat.Mode()), 8)),
		},
	}, nil
}

func parsePermissions(permissions string) (fs.FileMode, error) {
	// Check if there are 4 digits in the string.
	// Only 3 or 4 octal digits are allowed.
	if len(permissions) != 3 && len(permissions) != 4 {
		return 0, fmt.Errorf("permissions must be 3 or 4 octal digits")
	}

	// Check if there is a leading zero. If not exist then add 0 at the start of the string.
	// But only for 3 digits.
	if !strings.HasPrefix(permissions, "0") {
		if len(permissions) == 4 {
			return 0, fmt.Errorf("permissions must start with 0. Only files are supported")
		}
		permissions = "0" + permissions
	}

	parsed, err := strconv.ParseUint(permissions, 0, 32)
	if err != nil {
		return 0, fmt.Errorf("parsing permissions: %w. Please use UNIX style", err)
	}

	return fs.FileMode(parsed), nil
}

func strPtr(i string) *string {
	return &i
}
