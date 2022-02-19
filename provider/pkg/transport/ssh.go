package transport

import (
	"bytes"
	"context"
	"time"

	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/retry"
)

type SSHTransport struct {
	Addr         string
	clientConfig *ssh.ClientConfig
}

func (s *SSHTransport) dial(ctx context.Context, config *ssh.ClientConfig) (*ssh.Client, error) {
	var client *ssh.Client
	var err error
	_, _, err = retry.Until(ctx, retry.Acceptor{
		Accept: func(try int, nextRetryTime time.Duration) (bool, interface{}, error) {
			client, err = ssh.Dial("tcp", s.Addr, config)
			if err != nil {
				if try > 10 {
					return true, nil, err
				}
				return false, nil, nil
			}
			return true, nil, nil
		},
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func SSHInit(addr, user, key string) (*SSHTransport, error) {
	signer, err := ssh.ParsePrivateKey([]byte(key))
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}
	config.HostKeyCallback = ssh.InsecureIgnoreHostKey() //nolint:gosec

	s := &SSHTransport{
		Addr:         addr,
		clientConfig: config,
	}

	return s, nil
}

func (s *SSHTransport) RunCmd(ctx context.Context, cmd string) ([]byte, error) {
	client, err := s.dial(ctx, s.clientConfig)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var o bytes.Buffer
	session.Stdout = &o
	if err := session.Run(cmd); err != nil {
		return nil, err
	}

	return o.Bytes(), nil
}

func (s *SSHTransport) CopyFile(ctx context.Context, source, dest string) error {
	client, err := s.dial(ctx, s.clientConfig)
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	return scp.CopyPath(source, dest, session)
}
