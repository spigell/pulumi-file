package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/spigell/pulumi-file/sdk/go/file/remote"
	"os"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Read the private key from a file.
		key, err := os.ReadFile("../../misc/ssh/id_rsa")
		if err != nil {
			return err
		}

		file, err := remote.NewFile(ctx, "file", &remote.FileArgs{
			Connection: &remote.ConnectionArgs{
			     Host: pulumi.String("127.0.0.1"),
			     Port: pulumi.Int(2222),
			     User: pulumi.String("ssh-user"),
			     PrivateKey: pulumi.String(key),
			},
			UseSudo: pulumi.Bool(false),
			Permissions: pulumi.String("777"),
			Path: pulumi.String("pulumi.txt"),
			Content: pulumi.String("#Managed by Pulumi with file `resource` (GO)!!!!!!!!\n"),
		})
		if err != nil {
			return err
		}

		sudoFile, err := remote.NewFile(ctx, "sudo-file", &remote.FileArgs{
			Connection: &remote.ConnectionArgs{
			     Host: pulumi.String("127.0.0.1"),
			     Port: pulumi.Int(2222),
			     User: pulumi.String("ssh-user"),
			     PrivateKey: pulumi.String(key),
			},
			UseSudo: pulumi.Bool(true),
			Permissions: pulumi.String("644"),
			Path: pulumi.String("/etc/pulumi-sudo-user"),
			Content: pulumi.String("#Managed by Pulumi with file `resource` (GO)!!!!!!!!!!!!!!!\n"),
			Triggers: pulumi.Array{
				file.Md5sum,
				file.Permissions,
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("file", file.ID())
		ctx.Export("md5sum", file.Md5sum)

		ctx.Export("sudo-file", sudoFile.ID())
		ctx.Export("sudo-file.md5sum", sudoFile.Md5sum)
		return nil
	})
}
