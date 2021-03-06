// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package file

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Connection struct {
	Address    *string `pulumi:"address"`
	PrivateKey *string `pulumi:"privateKey"`
	User       *string `pulumi:"user"`
}

// ConnectionInput is an input type that accepts ConnectionArgs and ConnectionOutput values.
// You can construct a concrete instance of `ConnectionInput` via:
//
//          ConnectionArgs{...}
type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(context.Context) ConnectionOutput
}

type ConnectionArgs struct {
	Address    pulumi.StringPtrInput `pulumi:"address"`
	PrivateKey pulumi.StringPtrInput `pulumi:"privateKey"`
	User       pulumi.StringPtrInput `pulumi:"user"`
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil)).Elem()
}

func (i ConnectionArgs) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i ConnectionArgs) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

func (i ConnectionArgs) ToConnectionPtrOutput() ConnectionPtrOutput {
	return i.ToConnectionPtrOutputWithContext(context.Background())
}

func (i ConnectionArgs) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput).ToConnectionPtrOutputWithContext(ctx)
}

// ConnectionPtrInput is an input type that accepts ConnectionArgs, ConnectionPtr and ConnectionPtrOutput values.
// You can construct a concrete instance of `ConnectionPtrInput` via:
//
//          ConnectionArgs{...}
//
//  or:
//
//          nil
type ConnectionPtrInput interface {
	pulumi.Input

	ToConnectionPtrOutput() ConnectionPtrOutput
	ToConnectionPtrOutputWithContext(context.Context) ConnectionPtrOutput
}

type connectionPtrType ConnectionArgs

func ConnectionPtr(v *ConnectionArgs) ConnectionPtrInput {
	return (*connectionPtrType)(v)
}

func (*connectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *connectionPtrType) ToConnectionPtrOutput() ConnectionPtrOutput {
	return i.ToConnectionPtrOutputWithContext(context.Background())
}

func (i *connectionPtrType) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPtrOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionPtrOutput() ConnectionPtrOutput {
	return o.ToConnectionPtrOutputWithContext(context.Background())
}

func (o ConnectionOutput) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Connection) *Connection {
		return &v
	}).(ConnectionPtrOutput)
}

func (o ConnectionOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o ConnectionOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o ConnectionOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.User }).(pulumi.StringPtrOutput)
}

type ConnectionPtrOutput struct{ *pulumi.OutputState }

func (ConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionPtrOutput) ToConnectionPtrOutput() ConnectionPtrOutput {
	return o
}

func (o ConnectionPtrOutput) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return o
}

func (o ConnectionPtrOutput) Elem() ConnectionOutput {
	return o.ApplyT(func(v *Connection) Connection {
		if v != nil {
			return *v
		}
		var ret Connection
		return ret
	}).(ConnectionOutput)
}

func (o ConnectionPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPtrOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) *string {
		if v == nil {
			return nil
		}
		return v.PrivateKey
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPtrOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) *string {
		if v == nil {
			return nil
		}
		return v.User
	}).(pulumi.StringPtrOutput)
}

type Hooks struct {
	CommandAfterCreate  *string `pulumi:"commandAfterCreate"`
	CommandAfterDestroy *string `pulumi:"commandAfterDestroy"`
	CommandAfterUpdate  *string `pulumi:"commandAfterUpdate"`
}

// Defaults sets the appropriate defaults for Hooks
func (val *Hooks) Defaults() *Hooks {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CommandAfterCreate) {
		commandAfterCreate_ := ""
		tmp.CommandAfterCreate = &commandAfterCreate_
	}
	if isZero(tmp.CommandAfterDestroy) {
		commandAfterDestroy_ := ""
		tmp.CommandAfterDestroy = &commandAfterDestroy_
	}
	if isZero(tmp.CommandAfterUpdate) {
		commandAfterUpdate_ := ""
		tmp.CommandAfterUpdate = &commandAfterUpdate_
	}
	return &tmp
}

// HooksInput is an input type that accepts HooksArgs and HooksOutput values.
// You can construct a concrete instance of `HooksInput` via:
//
//          HooksArgs{...}
type HooksInput interface {
	pulumi.Input

	ToHooksOutput() HooksOutput
	ToHooksOutputWithContext(context.Context) HooksOutput
}

type HooksArgs struct {
	CommandAfterCreate  pulumi.StringPtrInput `pulumi:"commandAfterCreate"`
	CommandAfterDestroy pulumi.StringPtrInput `pulumi:"commandAfterDestroy"`
	CommandAfterUpdate  pulumi.StringPtrInput `pulumi:"commandAfterUpdate"`
}

func (HooksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Hooks)(nil)).Elem()
}

func (i HooksArgs) ToHooksOutput() HooksOutput {
	return i.ToHooksOutputWithContext(context.Background())
}

func (i HooksArgs) ToHooksOutputWithContext(ctx context.Context) HooksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HooksOutput)
}

func (i HooksArgs) ToHooksPtrOutput() HooksPtrOutput {
	return i.ToHooksPtrOutputWithContext(context.Background())
}

func (i HooksArgs) ToHooksPtrOutputWithContext(ctx context.Context) HooksPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HooksOutput).ToHooksPtrOutputWithContext(ctx)
}

// HooksPtrInput is an input type that accepts HooksArgs, HooksPtr and HooksPtrOutput values.
// You can construct a concrete instance of `HooksPtrInput` via:
//
//          HooksArgs{...}
//
//  or:
//
//          nil
type HooksPtrInput interface {
	pulumi.Input

	ToHooksPtrOutput() HooksPtrOutput
	ToHooksPtrOutputWithContext(context.Context) HooksPtrOutput
}

type hooksPtrType HooksArgs

func HooksPtr(v *HooksArgs) HooksPtrInput {
	return (*hooksPtrType)(v)
}

func (*hooksPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Hooks)(nil)).Elem()
}

func (i *hooksPtrType) ToHooksPtrOutput() HooksPtrOutput {
	return i.ToHooksPtrOutputWithContext(context.Background())
}

func (i *hooksPtrType) ToHooksPtrOutputWithContext(ctx context.Context) HooksPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HooksPtrOutput)
}

type HooksOutput struct{ *pulumi.OutputState }

func (HooksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Hooks)(nil)).Elem()
}

func (o HooksOutput) ToHooksOutput() HooksOutput {
	return o
}

func (o HooksOutput) ToHooksOutputWithContext(ctx context.Context) HooksOutput {
	return o
}

func (o HooksOutput) ToHooksPtrOutput() HooksPtrOutput {
	return o.ToHooksPtrOutputWithContext(context.Background())
}

func (o HooksOutput) ToHooksPtrOutputWithContext(ctx context.Context) HooksPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Hooks) *Hooks {
		return &v
	}).(HooksPtrOutput)
}

func (o HooksOutput) CommandAfterCreate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Hooks) *string { return v.CommandAfterCreate }).(pulumi.StringPtrOutput)
}

func (o HooksOutput) CommandAfterDestroy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Hooks) *string { return v.CommandAfterDestroy }).(pulumi.StringPtrOutput)
}

func (o HooksOutput) CommandAfterUpdate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Hooks) *string { return v.CommandAfterUpdate }).(pulumi.StringPtrOutput)
}

type HooksPtrOutput struct{ *pulumi.OutputState }

func (HooksPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Hooks)(nil)).Elem()
}

func (o HooksPtrOutput) ToHooksPtrOutput() HooksPtrOutput {
	return o
}

func (o HooksPtrOutput) ToHooksPtrOutputWithContext(ctx context.Context) HooksPtrOutput {
	return o
}

func (o HooksPtrOutput) Elem() HooksOutput {
	return o.ApplyT(func(v *Hooks) Hooks {
		if v != nil {
			return *v
		}
		var ret Hooks
		return ret
	}).(HooksOutput)
}

func (o HooksPtrOutput) CommandAfterCreate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hooks) *string {
		if v == nil {
			return nil
		}
		return v.CommandAfterCreate
	}).(pulumi.StringPtrOutput)
}

func (o HooksPtrOutput) CommandAfterDestroy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hooks) *string {
		if v == nil {
			return nil
		}
		return v.CommandAfterDestroy
	}).(pulumi.StringPtrOutput)
}

func (o HooksPtrOutput) CommandAfterUpdate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hooks) *string {
		if v == nil {
			return nil
		}
		return v.CommandAfterUpdate
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), ConnectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionPtrInput)(nil)).Elem(), ConnectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HooksInput)(nil)).Elem(), HooksArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HooksPtrInput)(nil)).Elem(), HooksArgs{})
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(ConnectionPtrOutput{})
	pulumi.RegisterOutputType(HooksOutput{})
	pulumi.RegisterOutputType(HooksPtrOutput{})
}
