# A file Pulumi Provider

This is a very simple pulumi native provider.

Resources:
* Node: It manages a content of file on linux machines. It supports only remote file now.

There is typescript and golang examples of using the single resource defined in `examples directory`.

``Read`` call is not implemented so commands like `pulumi refresh` or `pulumi import` are not working right now.

## Installing

This package is available only for JS/TS or Golang.

Grab plugin if pulumi didn't make it automaticaly:

```
pulumi plugin install resource file v0.0.2 --server https://github.com/spigell/pulumi-file/releases/download/v0.0.2
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

- Go 1.15
- NodeJS 10.X.X or later

### Local Build and Test

Most of the code for the provider implementation is in `provider/pkg` directory.  

A code generator is available which generates SDKs in TypeScript, Python, Go and .NET which are also checked in to the `sdk` folder.  The SDKs are generated from a schema in `schema.json`.  This file should be kept aligned with the resources, functions and types supported by the provider implementation.

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
