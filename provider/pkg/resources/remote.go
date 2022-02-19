package resources

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/spigell/pulumi-file/provider/pkg/transport"
)

func convertToRemote(r resource.PropertyMap) (*RemoteFile, error) {
	parsedPerm, err := strconv.ParseUint(r["permissions"].StringValue(), 0, 32)
	if err != nil {
		return &RemoteFile{}, fmt.Errorf("failed to convert permissions for file: %w", err)
	}

	n := &RemoteFile{
		Connection: Connection{
			Address:    secretOrPlain(r["connection"], resource.PropertyKey("address")),
			User:       secretOrPlain(r["connection"], resource.PropertyKey("user")),
			PrivateKey: secretOrPlain(r["connection"], resource.PropertyKey("privateKey")),
		},
		Content:     []byte(secretOrPlain(r["content"], resource.PropertyKey(""))),
		Permissions: fs.FileMode(parsedPerm),
		Path:        r["path"].StringValue(),
	}

	if raw, ok := r["hooks"]; ok {
		hooks := raw.ObjectValue()
		n.Hooks.CommandAfterCreate = valueOrEmpty(hooks, resource.PropertyKey("commandAfterCreate"))
		n.Hooks.CommandAfterUpdate = valueOrEmpty(hooks, resource.PropertyKey("commandAfterUpdate"))
		n.Hooks.CommandAfterDestroy = valueOrEmpty(hooks, resource.PropertyKey("commandAfterDestroy"))
	}

	n.WritableTempDirectory = valueOrEmpty(r, resource.PropertyKey("writableTempDirectory"))
	n.UseSudo = r["useSudo"].BoolValue()

	return n, nil
}

func ManageRemoteFile(ctx context.Context, kind string, m resource.PropertyMap) (map[string]interface{}, error) {
	r, err := convertToRemote(m)
	if err != nil {
		return nil, err
	}

	s, err := transport.SSHInit(r.Connection.Address, r.Connection.User, r.Connection.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed init ssh: %w", err)
	}

	r.transport = s

	if err != nil {
		return nil, err
	}

	var res map[string]interface{}
	switch kind {
	case "delete":
		res, err = r.delete(ctx)
		if err != nil {
			return nil, err
		}

	default:
		res, err = r.deploy(ctx, kind)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (r *RemoteFile) deploy(ctx context.Context, kind string) (map[string]interface{}, error) {
	tmp, err := ioutil.TempFile(os.TempDir(), "pulumi-file.*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	defer os.Remove(tmp.Name())

	err = ioutil.WriteFile(tmp.Name(), r.Content, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to write temp file: %w", err)
	}

	err = tmp.Chmod(r.Permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to change permissions for file: %w", err)
	}

	filename := filepath.Base(r.Path)

	var tmpOnRemote string

	if r.WritableTempDirectory != "" {
		tmpOnRemote = fmt.Sprintf("%s/%s", r.WritableTempDirectory, filename)
	} else {
		if r.Connection.User == "root" {
			tmpOnRemote = fmt.Sprintf("/root/%s", filename)
		} else {
			tmpOnRemote = fmt.Sprintf("/home/%s/%s", r.Connection.User, filename)
		}
	}

	err = r.transport.CopyFile(ctx, tmp.Name(), tmpOnRemote)
	if err != nil {
		return nil, fmt.Errorf("failed to copy temp file over scp to %s: %w", tmpOnRemote, err)
	}

	mvCmd := fmt.Sprintf("mv -v %s %s", tmpOnRemote, r.Path)

	if r.UseSudo {
		mvCmd = "sudo " + mvCmd
	}

	_, err = r.transport.RunCmd(ctx, mvCmd)

	if err != nil {
		return nil, fmt.Errorf("failed to execute move command %s: %w", mvCmd, err)
	}

	var hookCmd string
	switch kind {
	case "create":
		hookCmd = r.Hooks.CommandAfterCreate
	case "update":
		hookCmd = r.Hooks.CommandAfterUpdate
	default:
	}
	_, err = r.transport.RunCmd(ctx, hookCmd)
	if err != nil {
		return nil, fmt.Errorf("failed to execute hook command %s: %w", hookCmd, err)
	}

	result := make(map[string]interface{})
	result["id"] = r.Path
	return result, nil
}

func (r *RemoteFile) delete(ctx context.Context) (map[string]interface{}, error) {
	delCmd := fmt.Sprintf("rm -rfv %s", r.Path)

	if r.UseSudo {
		delCmd = "sudo " + delCmd
	}
	_, err := r.transport.RunCmd(ctx, delCmd)
	if err != nil {
		return nil, err
	}

	if r.Hooks.CommandAfterDestroy != "" {
		_, err = r.transport.RunCmd(ctx, r.Hooks.CommandAfterDestroy)
		if err != nil {
			return nil, fmt.Errorf("failed to execute post command %s: %w", r.Hooks.CommandAfterDestroy, err)
		}
	}
	result := make(map[string]interface{})
	return result, nil
}

func secretOrPlain(c resource.PropertyValue, token resource.PropertyKey) string {
	if c.IsSecret() {
		if token == "" {
			return c.SecretValue().Element.StringValue()
		}

		return c.SecretValue().Element.ObjectValue()[token].StringValue()
	}
	if token == "" {
		return c.StringValue()
	}

	if c.ObjectValue()[token].IsSecret() {
		return c.ObjectValue()[token].SecretValue().Element.StringValue()
	}

	return c.ObjectValue()[token].StringValue()
}

func valueOrEmpty(input resource.PropertyMap, tok resource.PropertyKey) string {
	if val, has := input[tok]; has {
		return val.StringValue()
	}
	return ""
}
