package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/spigell/pulumi-file/sdk/go/file"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
                key := cfg.RequireSecret("sshKey")

		file, err := file.NewRemote(ctx, "motd", &file.RemoteArgs{
			Connection: &file.ConnectionArgs{
			     Address: pulumi.String("127.0.0.1:2222"),
			     User: pulumi.String("ssh-user"),
			     PrivateKey: key,
			},
			Hooks: &file.HooksArgs{
				CommandAfterCreate: pulumi.String("sudo bash -c 'cat /etc/hello2.txt > /etc/motd'"),
				CommandAfterUpdate: pulumi.String("sudo bash -c 'cat /etc/hello2.txt > /etc/motd'"),
			},
			UseSudo: pulumi.Bool(true),
			WritableTempDirectory: pulumi.String("/config"),
			Permissions: pulumi.String("0644"),
			Path: pulumi.String("/etc/hello2.txt"),
			Content: pulumi.String("#Managed by Pulumi with file `resource` (GO)!\n"),
		})
		if err != nil {
			return err
		}

		ctx.Export("file", file.ID())
		return nil
	})
}
