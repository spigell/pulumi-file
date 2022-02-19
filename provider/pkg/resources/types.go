package resources

import (
	"io/fs"
	"github.com/spigell/pulumi-file/provider/pkg/transport"
)

type RemoteFile struct {
	Connection  Connection
	Content     []byte
	Permissions fs.FileMode
	WritableTempDirectory string
	UseSudo bool
	Path        string
	Hooks      Hooks
	transport *transport.SSHTransport
}

type Connection struct {
	Address    string
	User       string
	PrivateKey string
}

type Hooks struct {
	CommandAfterCreate  string
	CommandAfterUpdate  string
	CommandAfterDestroy string
}
