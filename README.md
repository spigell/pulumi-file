# A file Pulumi Provider

This is a very simple pulumi native provider.

Hugely based on [pulumi-go-provider](https://github.com/pulumi/pulumi-go-provider) and [pulumi-command-provider](https://github.com/pulumi/pulumi-command). Several pieces of code were from there.

Resources:
* Remote.File: It manages the content of files on Linux machines. It supports only remote files for now.

There are typescript and golang examples of using the single resource defined in `examples directory`.

``Read`` call is implemented so commands like `pulumi refresh` works now.
But `pulumi import` is not implemented now.

There are golang and typescript sdk. Another are generated but not uploaded to packages repositories.

## Installing

Grab the plugin if pulumi didn't make it automatically:

```
pulumi plugin install resource file v0.0.3 --server https://github.com/spigell/pulumi-file/releases/download/v0.0.3
```

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @spigell/pulumi-file

or `yarn`:

    $ yarn add @spigell/pulumi-file

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/spigell/pulumi-file/sdk

## Building

### Dependencies

- Go 1.21
- NodeJS 16.X.X or later

### Local Build and Test

Most of the code for the provider implementation is in `provider/pkg` directory.  

A code generator is available which generates SDKs in TypeScript, Python, Go and .NET which are also checked in to the `sdk` folder.  The SDKs are generated from a schema in `schema.json`.

Note that the generated provider plugin (`pulumi-resource-file`) must be on your `PATH` to be used by Pulumi deployments.

```bash
# build the plugin and provider
$ VERSION=<VERSION> make build

# test
$ cd examples/simple-go
$ go mod tidy
$ pulumi stack init test
$ PATH=$PATH:../../bin pulumi up
```
