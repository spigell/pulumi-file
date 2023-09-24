// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package remote

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/spigell/pulumi-file/sdk/go/file/internal"
)

var _ = internal.GetEnvOrDefault

// Instructions for how to connect to a remote endpoint.
type Connection struct {
	// SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
	AgentSocketPath *string `pulumi:"agentSocketPath"`
	// Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
	DialErrorLimit *int `pulumi:"dialErrorLimit"`
	// The address of the resource to connect to.
	Host string `pulumi:"host"`
	// The password we should use for the connection.
	Password *string `pulumi:"password"`
	// Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
	PerDialTimeout *int `pulumi:"perDialTimeout"`
	// The port to connect to.
	Port *int `pulumi:"port"`
	// The contents of an SSH key to use for the connection. This takes preference over the password if provided.
	PrivateKey *string `pulumi:"privateKey"`
	// The password to use in case the private key is encrypted.
	PrivateKeyPassword *string `pulumi:"privateKeyPassword"`
	// The connection settings for the bastion/proxy host.
	Proxy *ProxyConnection `pulumi:"proxy"`
	// The user that we should use for the connection.
	User *string `pulumi:"user"`
}

// Defaults sets the appropriate defaults for Connection
func (val *Connection) Defaults() *Connection {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.DialErrorLimit == nil {
		dialErrorLimit_ := 10
		tmp.DialErrorLimit = &dialErrorLimit_
	}
	if tmp.PerDialTimeout == nil {
		perDialTimeout_ := 15
		tmp.PerDialTimeout = &perDialTimeout_
	}
	if tmp.Port == nil {
		port_ := 22
		tmp.Port = &port_
	}
	tmp.Proxy = tmp.Proxy.Defaults()

	if tmp.User == nil {
		user_ := "root"
		tmp.User = &user_
	}
	return &tmp
}

// ConnectionInput is an input type that accepts ConnectionArgs and ConnectionOutput values.
// You can construct a concrete instance of `ConnectionInput` via:
//
//	ConnectionArgs{...}
type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(context.Context) ConnectionOutput
}

// Instructions for how to connect to a remote endpoint.
type ConnectionArgs struct {
	// SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
	AgentSocketPath pulumi.StringPtrInput `pulumi:"agentSocketPath"`
	// Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
	DialErrorLimit pulumi.IntPtrInput `pulumi:"dialErrorLimit"`
	// The address of the resource to connect to.
	Host pulumi.StringInput `pulumi:"host"`
	// The password we should use for the connection.
	Password pulumi.StringPtrInput `pulumi:"password"`
	// Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
	PerDialTimeout pulumi.IntPtrInput `pulumi:"perDialTimeout"`
	// The port to connect to.
	Port pulumi.IntPtrInput `pulumi:"port"`
	// The contents of an SSH key to use for the connection. This takes preference over the password if provided.
	PrivateKey pulumi.StringPtrInput `pulumi:"privateKey"`
	// The password to use in case the private key is encrypted.
	PrivateKeyPassword pulumi.StringPtrInput `pulumi:"privateKeyPassword"`
	// The connection settings for the bastion/proxy host.
	Proxy ProxyConnectionPtrInput `pulumi:"proxy"`
	// The user that we should use for the connection.
	User pulumi.StringPtrInput `pulumi:"user"`
}

// Defaults sets the appropriate defaults for ConnectionArgs
func (val *ConnectionArgs) Defaults() *ConnectionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.DialErrorLimit == nil {
		tmp.DialErrorLimit = pulumi.IntPtr(10)
	}
	if tmp.PerDialTimeout == nil {
		tmp.PerDialTimeout = pulumi.IntPtr(15)
	}
	if tmp.Port == nil {
		tmp.Port = pulumi.IntPtr(22)
	}

	if tmp.User == nil {
		tmp.User = pulumi.StringPtr("root")
	}
	return &tmp
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

func (i ConnectionArgs) ToOutput(ctx context.Context) pulumix.Output[Connection] {
	return pulumix.Output[Connection]{
		OutputState: i.ToConnectionOutputWithContext(ctx).OutputState,
	}
}

// Instructions for how to connect to a remote endpoint.
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

func (o ConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[Connection] {
	return pulumix.Output[Connection]{
		OutputState: o.OutputState,
	}
}

// SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
func (o ConnectionOutput) AgentSocketPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.AgentSocketPath }).(pulumi.StringPtrOutput)
}

// Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
func (o ConnectionOutput) DialErrorLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Connection) *int { return v.DialErrorLimit }).(pulumi.IntPtrOutput)
}

// The address of the resource to connect to.
func (o ConnectionOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v Connection) string { return v.Host }).(pulumi.StringOutput)
}

// The password we should use for the connection.
func (o ConnectionOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.Password }).(pulumi.StringPtrOutput)
}

// Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
func (o ConnectionOutput) PerDialTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Connection) *int { return v.PerDialTimeout }).(pulumi.IntPtrOutput)
}

// The port to connect to.
func (o ConnectionOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Connection) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// The contents of an SSH key to use for the connection. This takes preference over the password if provided.
func (o ConnectionOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

// The password to use in case the private key is encrypted.
func (o ConnectionOutput) PrivateKeyPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.PrivateKeyPassword }).(pulumi.StringPtrOutput)
}

// The connection settings for the bastion/proxy host.
func (o ConnectionOutput) Proxy() ProxyConnectionPtrOutput {
	return o.ApplyT(func(v Connection) *ProxyConnection { return v.Proxy }).(ProxyConnectionPtrOutput)
}

// The user that we should use for the connection.
func (o ConnectionOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Connection) *string { return v.User }).(pulumi.StringPtrOutput)
}

// Instructions for how to connect to a remote endpoint via a bastion host.
type ProxyConnection struct {
	// SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
	AgentSocketPath *string `pulumi:"agentSocketPath"`
	// Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
	DialErrorLimit *int `pulumi:"dialErrorLimit"`
	// The address of the bastion host to connect to.
	Host string `pulumi:"host"`
	// The password we should use for the connection to the bastion host.
	Password *string `pulumi:"password"`
	// Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
	PerDialTimeout *int `pulumi:"perDialTimeout"`
	// The port of the bastion host to connect to.
	Port *int `pulumi:"port"`
	// The contents of an SSH key to use for the connection. This takes preference over the password if provided.
	PrivateKey *string `pulumi:"privateKey"`
	// The password to use in case the private key is encrypted.
	PrivateKeyPassword *string `pulumi:"privateKeyPassword"`
	// The user that we should use for the connection to the bastion host.
	User *string `pulumi:"user"`
}

// Defaults sets the appropriate defaults for ProxyConnection
func (val *ProxyConnection) Defaults() *ProxyConnection {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.DialErrorLimit == nil {
		dialErrorLimit_ := 10
		tmp.DialErrorLimit = &dialErrorLimit_
	}
	if tmp.PerDialTimeout == nil {
		perDialTimeout_ := 15
		tmp.PerDialTimeout = &perDialTimeout_
	}
	if tmp.Port == nil {
		port_ := 22
		tmp.Port = &port_
	}
	if tmp.User == nil {
		user_ := "root"
		tmp.User = &user_
	}
	return &tmp
}

// ProxyConnectionInput is an input type that accepts ProxyConnectionArgs and ProxyConnectionOutput values.
// You can construct a concrete instance of `ProxyConnectionInput` via:
//
//	ProxyConnectionArgs{...}
type ProxyConnectionInput interface {
	pulumi.Input

	ToProxyConnectionOutput() ProxyConnectionOutput
	ToProxyConnectionOutputWithContext(context.Context) ProxyConnectionOutput
}

// Instructions for how to connect to a remote endpoint via a bastion host.
type ProxyConnectionArgs struct {
	// SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
	AgentSocketPath pulumi.StringPtrInput `pulumi:"agentSocketPath"`
	// Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
	DialErrorLimit pulumi.IntPtrInput `pulumi:"dialErrorLimit"`
	// The address of the bastion host to connect to.
	Host pulumi.StringInput `pulumi:"host"`
	// The password we should use for the connection to the bastion host.
	Password pulumi.StringPtrInput `pulumi:"password"`
	// Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
	PerDialTimeout pulumi.IntPtrInput `pulumi:"perDialTimeout"`
	// The port of the bastion host to connect to.
	Port pulumi.IntPtrInput `pulumi:"port"`
	// The contents of an SSH key to use for the connection. This takes preference over the password if provided.
	PrivateKey pulumi.StringPtrInput `pulumi:"privateKey"`
	// The password to use in case the private key is encrypted.
	PrivateKeyPassword pulumi.StringPtrInput `pulumi:"privateKeyPassword"`
	// The user that we should use for the connection to the bastion host.
	User pulumi.StringPtrInput `pulumi:"user"`
}

// Defaults sets the appropriate defaults for ProxyConnectionArgs
func (val *ProxyConnectionArgs) Defaults() *ProxyConnectionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.DialErrorLimit == nil {
		tmp.DialErrorLimit = pulumi.IntPtr(10)
	}
	if tmp.PerDialTimeout == nil {
		tmp.PerDialTimeout = pulumi.IntPtr(15)
	}
	if tmp.Port == nil {
		tmp.Port = pulumi.IntPtr(22)
	}
	if tmp.User == nil {
		tmp.User = pulumi.StringPtr("root")
	}
	return &tmp
}
func (ProxyConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProxyConnection)(nil)).Elem()
}

func (i ProxyConnectionArgs) ToProxyConnectionOutput() ProxyConnectionOutput {
	return i.ToProxyConnectionOutputWithContext(context.Background())
}

func (i ProxyConnectionArgs) ToProxyConnectionOutputWithContext(ctx context.Context) ProxyConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyConnectionOutput)
}

func (i ProxyConnectionArgs) ToOutput(ctx context.Context) pulumix.Output[ProxyConnection] {
	return pulumix.Output[ProxyConnection]{
		OutputState: i.ToProxyConnectionOutputWithContext(ctx).OutputState,
	}
}

func (i ProxyConnectionArgs) ToProxyConnectionPtrOutput() ProxyConnectionPtrOutput {
	return i.ToProxyConnectionPtrOutputWithContext(context.Background())
}

func (i ProxyConnectionArgs) ToProxyConnectionPtrOutputWithContext(ctx context.Context) ProxyConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyConnectionOutput).ToProxyConnectionPtrOutputWithContext(ctx)
}

// ProxyConnectionPtrInput is an input type that accepts ProxyConnectionArgs, ProxyConnectionPtr and ProxyConnectionPtrOutput values.
// You can construct a concrete instance of `ProxyConnectionPtrInput` via:
//
//	        ProxyConnectionArgs{...}
//
//	or:
//
//	        nil
type ProxyConnectionPtrInput interface {
	pulumi.Input

	ToProxyConnectionPtrOutput() ProxyConnectionPtrOutput
	ToProxyConnectionPtrOutputWithContext(context.Context) ProxyConnectionPtrOutput
}

type proxyConnectionPtrType ProxyConnectionArgs

func ProxyConnectionPtr(v *ProxyConnectionArgs) ProxyConnectionPtrInput {
	return (*proxyConnectionPtrType)(v)
}

func (*proxyConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyConnection)(nil)).Elem()
}

func (i *proxyConnectionPtrType) ToProxyConnectionPtrOutput() ProxyConnectionPtrOutput {
	return i.ToProxyConnectionPtrOutputWithContext(context.Background())
}

func (i *proxyConnectionPtrType) ToProxyConnectionPtrOutputWithContext(ctx context.Context) ProxyConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyConnectionPtrOutput)
}

func (i *proxyConnectionPtrType) ToOutput(ctx context.Context) pulumix.Output[*ProxyConnection] {
	return pulumix.Output[*ProxyConnection]{
		OutputState: i.ToProxyConnectionPtrOutputWithContext(ctx).OutputState,
	}
}

// Instructions for how to connect to a remote endpoint via a bastion host.
type ProxyConnectionOutput struct{ *pulumi.OutputState }

func (ProxyConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProxyConnection)(nil)).Elem()
}

func (o ProxyConnectionOutput) ToProxyConnectionOutput() ProxyConnectionOutput {
	return o
}

func (o ProxyConnectionOutput) ToProxyConnectionOutputWithContext(ctx context.Context) ProxyConnectionOutput {
	return o
}

func (o ProxyConnectionOutput) ToProxyConnectionPtrOutput() ProxyConnectionPtrOutput {
	return o.ToProxyConnectionPtrOutputWithContext(context.Background())
}

func (o ProxyConnectionOutput) ToProxyConnectionPtrOutputWithContext(ctx context.Context) ProxyConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProxyConnection) *ProxyConnection {
		return &v
	}).(ProxyConnectionPtrOutput)
}

func (o ProxyConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[ProxyConnection] {
	return pulumix.Output[ProxyConnection]{
		OutputState: o.OutputState,
	}
}

// SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
func (o ProxyConnectionOutput) AgentSocketPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProxyConnection) *string { return v.AgentSocketPath }).(pulumi.StringPtrOutput)
}

// Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
func (o ProxyConnectionOutput) DialErrorLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProxyConnection) *int { return v.DialErrorLimit }).(pulumi.IntPtrOutput)
}

// The address of the bastion host to connect to.
func (o ProxyConnectionOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v ProxyConnection) string { return v.Host }).(pulumi.StringOutput)
}

// The password we should use for the connection to the bastion host.
func (o ProxyConnectionOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProxyConnection) *string { return v.Password }).(pulumi.StringPtrOutput)
}

// Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
func (o ProxyConnectionOutput) PerDialTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProxyConnection) *int { return v.PerDialTimeout }).(pulumi.IntPtrOutput)
}

// The port of the bastion host to connect to.
func (o ProxyConnectionOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProxyConnection) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// The contents of an SSH key to use for the connection. This takes preference over the password if provided.
func (o ProxyConnectionOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProxyConnection) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

// The password to use in case the private key is encrypted.
func (o ProxyConnectionOutput) PrivateKeyPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProxyConnection) *string { return v.PrivateKeyPassword }).(pulumi.StringPtrOutput)
}

// The user that we should use for the connection to the bastion host.
func (o ProxyConnectionOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProxyConnection) *string { return v.User }).(pulumi.StringPtrOutput)
}

type ProxyConnectionPtrOutput struct{ *pulumi.OutputState }

func (ProxyConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyConnection)(nil)).Elem()
}

func (o ProxyConnectionPtrOutput) ToProxyConnectionPtrOutput() ProxyConnectionPtrOutput {
	return o
}

func (o ProxyConnectionPtrOutput) ToProxyConnectionPtrOutputWithContext(ctx context.Context) ProxyConnectionPtrOutput {
	return o
}

func (o ProxyConnectionPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ProxyConnection] {
	return pulumix.Output[*ProxyConnection]{
		OutputState: o.OutputState,
	}
}

func (o ProxyConnectionPtrOutput) Elem() ProxyConnectionOutput {
	return o.ApplyT(func(v *ProxyConnection) ProxyConnection {
		if v != nil {
			return *v
		}
		var ret ProxyConnection
		return ret
	}).(ProxyConnectionOutput)
}

// SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
func (o ProxyConnectionPtrOutput) AgentSocketPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyConnection) *string {
		if v == nil {
			return nil
		}
		return v.AgentSocketPath
	}).(pulumi.StringPtrOutput)
}

// Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
func (o ProxyConnectionPtrOutput) DialErrorLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProxyConnection) *int {
		if v == nil {
			return nil
		}
		return v.DialErrorLimit
	}).(pulumi.IntPtrOutput)
}

// The address of the bastion host to connect to.
func (o ProxyConnectionPtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyConnection) *string {
		if v == nil {
			return nil
		}
		return &v.Host
	}).(pulumi.StringPtrOutput)
}

// The password we should use for the connection to the bastion host.
func (o ProxyConnectionPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyConnection) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

// Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
func (o ProxyConnectionPtrOutput) PerDialTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProxyConnection) *int {
		if v == nil {
			return nil
		}
		return v.PerDialTimeout
	}).(pulumi.IntPtrOutput)
}

// The port of the bastion host to connect to.
func (o ProxyConnectionPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProxyConnection) *int {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.IntPtrOutput)
}

// The contents of an SSH key to use for the connection. This takes preference over the password if provided.
func (o ProxyConnectionPtrOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyConnection) *string {
		if v == nil {
			return nil
		}
		return v.PrivateKey
	}).(pulumi.StringPtrOutput)
}

// The password to use in case the private key is encrypted.
func (o ProxyConnectionPtrOutput) PrivateKeyPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyConnection) *string {
		if v == nil {
			return nil
		}
		return v.PrivateKeyPassword
	}).(pulumi.StringPtrOutput)
}

// The user that we should use for the connection to the bastion host.
func (o ProxyConnectionPtrOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyConnection) *string {
		if v == nil {
			return nil
		}
		return v.User
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), ConnectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyConnectionInput)(nil)).Elem(), ProxyConnectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyConnectionPtrInput)(nil)).Elem(), ProxyConnectionArgs{})
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(ProxyConnectionOutput{})
	pulumi.RegisterOutputType(ProxyConnectionPtrOutput{})
}