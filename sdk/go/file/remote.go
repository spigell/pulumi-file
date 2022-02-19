// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package file

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Remote struct {
	pulumi.CustomResourceState

	Connection  ConnectionPtrOutput    `pulumi:"connection"`
	Content     pulumi.StringPtrOutput `pulumi:"content"`
	Path        pulumi.StringPtrOutput `pulumi:"path"`
	Permissions pulumi.StringPtrOutput `pulumi:"permissions"`
}

// NewRemote registers a new resource with the given unique name, arguments, and options.
func NewRemote(ctx *pulumi.Context,
	name string, args *RemoteArgs, opts ...pulumi.ResourceOption) (*Remote, error) {
	if args == nil {
		args = &RemoteArgs{}
	}

	if isZero(args.UseSudo) {
		args.UseSudo = pulumi.BoolPtr(false)
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"connection.address",
		"path",
		"permissions",
	})
	opts = append(opts, replaceOnChanges)
	opts = pkgResourceDefaultOpts(opts)
	var resource Remote
	err := ctx.RegisterResource("file:index:Remote", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemote gets an existing Remote resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemote(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemoteState, opts ...pulumi.ResourceOption) (*Remote, error) {
	var resource Remote
	err := ctx.ReadResource("file:index:Remote", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Remote resources.
type remoteState struct {
}

type RemoteState struct {
}

func (RemoteState) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteState)(nil)).Elem()
}

type remoteArgs struct {
	Connection            *Connection `pulumi:"connection"`
	Content               *string     `pulumi:"content"`
	Hooks                 *Hooks      `pulumi:"hooks"`
	Path                  *string     `pulumi:"path"`
	Permissions           *string     `pulumi:"permissions"`
	UseSudo               *bool       `pulumi:"useSudo"`
	WritableTempDirectory *string     `pulumi:"writableTempDirectory"`
}

// The set of arguments for constructing a Remote resource.
type RemoteArgs struct {
	Connection            ConnectionPtrInput
	Content               pulumi.StringPtrInput
	Hooks                 HooksPtrInput
	Path                  pulumi.StringPtrInput
	Permissions           pulumi.StringPtrInput
	UseSudo               pulumi.BoolPtrInput
	WritableTempDirectory pulumi.StringPtrInput
}

func (RemoteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteArgs)(nil)).Elem()
}

type RemoteInput interface {
	pulumi.Input

	ToRemoteOutput() RemoteOutput
	ToRemoteOutputWithContext(ctx context.Context) RemoteOutput
}

func (*Remote) ElementType() reflect.Type {
	return reflect.TypeOf((**Remote)(nil)).Elem()
}

func (i *Remote) ToRemoteOutput() RemoteOutput {
	return i.ToRemoteOutputWithContext(context.Background())
}

func (i *Remote) ToRemoteOutputWithContext(ctx context.Context) RemoteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteOutput)
}

// RemoteArrayInput is an input type that accepts RemoteArray and RemoteArrayOutput values.
// You can construct a concrete instance of `RemoteArrayInput` via:
//
//          RemoteArray{ RemoteArgs{...} }
type RemoteArrayInput interface {
	pulumi.Input

	ToRemoteArrayOutput() RemoteArrayOutput
	ToRemoteArrayOutputWithContext(context.Context) RemoteArrayOutput
}

type RemoteArray []RemoteInput

func (RemoteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Remote)(nil)).Elem()
}

func (i RemoteArray) ToRemoteArrayOutput() RemoteArrayOutput {
	return i.ToRemoteArrayOutputWithContext(context.Background())
}

func (i RemoteArray) ToRemoteArrayOutputWithContext(ctx context.Context) RemoteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteArrayOutput)
}

// RemoteMapInput is an input type that accepts RemoteMap and RemoteMapOutput values.
// You can construct a concrete instance of `RemoteMapInput` via:
//
//          RemoteMap{ "key": RemoteArgs{...} }
type RemoteMapInput interface {
	pulumi.Input

	ToRemoteMapOutput() RemoteMapOutput
	ToRemoteMapOutputWithContext(context.Context) RemoteMapOutput
}

type RemoteMap map[string]RemoteInput

func (RemoteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Remote)(nil)).Elem()
}

func (i RemoteMap) ToRemoteMapOutput() RemoteMapOutput {
	return i.ToRemoteMapOutputWithContext(context.Background())
}

func (i RemoteMap) ToRemoteMapOutputWithContext(ctx context.Context) RemoteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteMapOutput)
}

type RemoteOutput struct{ *pulumi.OutputState }

func (RemoteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Remote)(nil)).Elem()
}

func (o RemoteOutput) ToRemoteOutput() RemoteOutput {
	return o
}

func (o RemoteOutput) ToRemoteOutputWithContext(ctx context.Context) RemoteOutput {
	return o
}

type RemoteArrayOutput struct{ *pulumi.OutputState }

func (RemoteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Remote)(nil)).Elem()
}

func (o RemoteArrayOutput) ToRemoteArrayOutput() RemoteArrayOutput {
	return o
}

func (o RemoteArrayOutput) ToRemoteArrayOutputWithContext(ctx context.Context) RemoteArrayOutput {
	return o
}

func (o RemoteArrayOutput) Index(i pulumi.IntInput) RemoteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Remote {
		return vs[0].([]*Remote)[vs[1].(int)]
	}).(RemoteOutput)
}

type RemoteMapOutput struct{ *pulumi.OutputState }

func (RemoteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Remote)(nil)).Elem()
}

func (o RemoteMapOutput) ToRemoteMapOutput() RemoteMapOutput {
	return o
}

func (o RemoteMapOutput) ToRemoteMapOutputWithContext(ctx context.Context) RemoteMapOutput {
	return o
}

func (o RemoteMapOutput) MapIndex(k pulumi.StringInput) RemoteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Remote {
		return vs[0].(map[string]*Remote)[vs[1].(string)]
	}).(RemoteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteInput)(nil)).Elem(), &Remote{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteArrayInput)(nil)).Elem(), RemoteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteMapInput)(nil)).Elem(), RemoteMap{})
	pulumi.RegisterOutputType(RemoteOutput{})
	pulumi.RegisterOutputType(RemoteArrayOutput{})
	pulumi.RegisterOutputType(RemoteMapOutput{})
}
