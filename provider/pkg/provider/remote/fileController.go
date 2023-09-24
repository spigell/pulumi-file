package remote

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// These are not required. They indicate to Go that implements the following interfaces.
// If the function signature doesn't match or isn't implemented, we get nice compile time errors in this file.
var (
	_ = (infer.CustomResource[FileInputs, FileOutputs])((*File)(nil))
	_ = (infer.CustomUpdate[FileInputs, FileOutputs])((*File)(nil))
	_ = (infer.CustomDelete[FileOutputs])((*File)(nil))
	_ = (infer.CustomRead[FileInputs, FileOutputs])((*File)(nil))
)

// This is the Create method. This will be run on every resource creation.
func (*File) Create(ctx p.Context, name string, input FileInputs, preview bool) (string, FileOutputs, error) {
	state := FileOutputs{FileInputs: input}
	var err error
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return id, state, nil
	}

	if !preview {
		err = state.deploy(ctx, input)
	}
	return id, state, err
}

// The Update method will be run on every update.
func (*File) Update(ctx p.Context, _ string, olds FileOutputs, news FileInputs, preview bool) (FileOutputs, error) {
	state := FileOutputs{FileInputs: news, BaseOutputs: olds.BaseOutputs}
	if preview {
		return state, nil
	}
	var err error
	if !preview {
		err = state.deploy(ctx, news)
	}
	return state, err
}

// The Delete method will run when the resource is deleted.
func (*File) Delete(ctx p.Context, _ string, props FileOutputs) error {
	return props.delete(ctx, props)
}

func (*File) Read(ctx p.Context, id string, input FileInputs, output FileOutputs) (string, FileInputs, FileOutputs, error) {
	// Import is not yet implemented
	// Empty connection means we are importing
	if input.Connection == nil {
		return "", input, output, status.Error(codes.Unimplemented, "Import is not yet implemented")
	}

	read, err := output.read(ctx, input)
	if err != nil {
		return "", input, output, err
	}

	return id, input, *read, nil
}
