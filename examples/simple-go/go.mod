module wgpeer_local

go 1.14

require (
	github.com/pulumi/pulumi/sdk/v3 v3.24.1
	github.com/spigell/pulumi-file/sdk v0.0.0-00010101000000-000000000000
)

replace github.com/spigell/pulumi-file/sdk => ../../sdk
