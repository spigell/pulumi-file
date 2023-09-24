package remote

import (
	"github.com/pulumi/pulumi-go-provider/infer"
)

type File struct{}

// Implementing Annotate lets you provide descriptions for resources and they will
// be visible in the provider's schema and the generated SDKs.
func (f *File) Annotate(a infer.Annotator) {
	a.Describe(&f, `A file to be created on a remote host. The connection is established via ssh.`)
}

// The arguments for a remote Command resource.
type FileInputs struct {
	// the pulumi-go-provider library uses field tags to dictate behavior.
	// pulumi:"connection" specifies the name of the field in the schema
	// pulumi:"optional" specifies that a field is optional. This must be a pointer.
	// provider:"secret" specifies that a field should be marked secret.
	Connection  *Connection `pulumi:"connection" provider:"secret"`
	Triggers    *[]any      `pulumi:"triggers,optional"`
	Path        *string     `pulumi:"path"`
	Content     *string     `pulumi:"content,optional"`
	UseSudo     *bool       `pulumi:"useSudo,optional" provider:"replace_on_change"`
	SFTPPath    *string     `pulumi:"sftpPath,optional"`
	Permissions *string     `pulumi:"permissions,optional"`
}

// Implementing Annotate lets you provide descriptions and default values for arguments and they will
// be visible in the provider's schema and the generated SDKs.
func (f *FileInputs) Annotate(a infer.Annotator) {
	a.Describe(&f.Connection, "The parameters with which to connect to the remote host.")
	a.Describe(&f.Triggers, "Trigger replacements on changes to this input.")
	a.Describe(&f.Path, "The path for file on remote server")
	a.Describe(&f.UseSudo, "if enabled then use sudo for copy command instead of direct copy")
	a.Describe(&f.SFTPPath, "sudo mode requires a external sftp server to be running on remote host")
	a.SetDefault(&f.SFTPPath, "/usr/lib/ssh/sftp-server")
	a.Describe(&f.Content, "The content of file")
	a.Describe(&f.Permissions, "Unix permissions for file")
	a.SetDefault(&f.Permissions, "0664")
}

// The properties for a remote Command resource.
type FileOutputs struct {
	FileInputs
	BaseOutputs
}
