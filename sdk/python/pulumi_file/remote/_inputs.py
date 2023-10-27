# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ConnectionArgs',
    'ProxyConnectionArgs',
]

@pulumi.input_type
class ConnectionArgs:
    def __init__(__self__, *,
                 host: pulumi.Input[str],
                 agent_socket_path: Optional[pulumi.Input[str]] = None,
                 dial_error_limit: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 per_dial_timeout: Optional[pulumi.Input[int]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_password: Optional[pulumi.Input[str]] = None,
                 proxy: Optional[pulumi.Input['ProxyConnectionArgs']] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Instructions for how to connect to a remote endpoint.
        :param pulumi.Input[str] host: The address of the resource to connect to.
        :param pulumi.Input[str] agent_socket_path: SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
        :param pulumi.Input[int] dial_error_limit: Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
        :param pulumi.Input[str] password: The password we should use for the connection.
        :param pulumi.Input[int] per_dial_timeout: Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
        :param pulumi.Input[int] port: The port to connect to.
        :param pulumi.Input[str] private_key: The contents of an SSH key to use for the connection. This takes preference over the password if provided.
        :param pulumi.Input[str] private_key_password: The password to use in case the private key is encrypted.
        :param pulumi.Input['ProxyConnectionArgs'] proxy: The connection settings for the bastion/proxy host.
        :param pulumi.Input[str] user: The user that we should use for the connection.
        """
        ConnectionArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            host=host,
            agent_socket_path=agent_socket_path,
            dial_error_limit=dial_error_limit,
            password=password,
            per_dial_timeout=per_dial_timeout,
            port=port,
            private_key=private_key,
            private_key_password=private_key_password,
            proxy=proxy,
            user=user,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             host: Optional[pulumi.Input[str]] = None,
             agent_socket_path: Optional[pulumi.Input[str]] = None,
             dial_error_limit: Optional[pulumi.Input[int]] = None,
             password: Optional[pulumi.Input[str]] = None,
             per_dial_timeout: Optional[pulumi.Input[int]] = None,
             port: Optional[pulumi.Input[int]] = None,
             private_key: Optional[pulumi.Input[str]] = None,
             private_key_password: Optional[pulumi.Input[str]] = None,
             proxy: Optional[pulumi.Input['ProxyConnectionArgs']] = None,
             user: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if host is None:
            raise TypeError("Missing 'host' argument")
        if agent_socket_path is None and 'agentSocketPath' in kwargs:
            agent_socket_path = kwargs['agentSocketPath']
        if dial_error_limit is None and 'dialErrorLimit' in kwargs:
            dial_error_limit = kwargs['dialErrorLimit']
        if per_dial_timeout is None and 'perDialTimeout' in kwargs:
            per_dial_timeout = kwargs['perDialTimeout']
        if private_key is None and 'privateKey' in kwargs:
            private_key = kwargs['privateKey']
        if private_key_password is None and 'privateKeyPassword' in kwargs:
            private_key_password = kwargs['privateKeyPassword']

        _setter("host", host)
        if agent_socket_path is not None:
            _setter("agent_socket_path", agent_socket_path)
        if dial_error_limit is None:
            dial_error_limit = 10
        if dial_error_limit is not None:
            _setter("dial_error_limit", dial_error_limit)
        if password is not None:
            _setter("password", password)
        if per_dial_timeout is None:
            per_dial_timeout = 15
        if per_dial_timeout is not None:
            _setter("per_dial_timeout", per_dial_timeout)
        if port is None:
            port = 22
        if port is not None:
            _setter("port", port)
        if private_key is not None:
            _setter("private_key", private_key)
        if private_key_password is not None:
            _setter("private_key_password", private_key_password)
        if proxy is not None:
            _setter("proxy", proxy)
        if user is None:
            user = 'root'
        if user is not None:
            _setter("user", user)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Input[str]:
        """
        The address of the resource to connect to.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: pulumi.Input[str]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="agentSocketPath")
    def agent_socket_path(self) -> Optional[pulumi.Input[str]]:
        """
        SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
        """
        return pulumi.get(self, "agent_socket_path")

    @agent_socket_path.setter
    def agent_socket_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "agent_socket_path", value)

    @property
    @pulumi.getter(name="dialErrorLimit")
    def dial_error_limit(self) -> Optional[pulumi.Input[int]]:
        """
        Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
        """
        return pulumi.get(self, "dial_error_limit")

    @dial_error_limit.setter
    def dial_error_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dial_error_limit", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password we should use for the connection.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="perDialTimeout")
    def per_dial_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
        """
        return pulumi.get(self, "per_dial_timeout")

    @per_dial_timeout.setter
    def per_dial_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "per_dial_timeout", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port to connect to.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        The contents of an SSH key to use for the connection. This takes preference over the password if provided.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="privateKeyPassword")
    def private_key_password(self) -> Optional[pulumi.Input[str]]:
        """
        The password to use in case the private key is encrypted.
        """
        return pulumi.get(self, "private_key_password")

    @private_key_password.setter
    def private_key_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_password", value)

    @property
    @pulumi.getter
    def proxy(self) -> Optional[pulumi.Input['ProxyConnectionArgs']]:
        """
        The connection settings for the bastion/proxy host.
        """
        return pulumi.get(self, "proxy")

    @proxy.setter
    def proxy(self, value: Optional[pulumi.Input['ProxyConnectionArgs']]):
        pulumi.set(self, "proxy", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        The user that we should use for the connection.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


@pulumi.input_type
class ProxyConnectionArgs:
    def __init__(__self__, *,
                 host: pulumi.Input[str],
                 agent_socket_path: Optional[pulumi.Input[str]] = None,
                 dial_error_limit: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 per_dial_timeout: Optional[pulumi.Input[int]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_password: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Instructions for how to connect to a remote endpoint via a bastion host.
        :param pulumi.Input[str] host: The address of the bastion host to connect to.
        :param pulumi.Input[str] agent_socket_path: SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
        :param pulumi.Input[int] dial_error_limit: Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
        :param pulumi.Input[str] password: The password we should use for the connection to the bastion host.
        :param pulumi.Input[int] per_dial_timeout: Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
        :param pulumi.Input[int] port: The port of the bastion host to connect to.
        :param pulumi.Input[str] private_key: The contents of an SSH key to use for the connection. This takes preference over the password if provided.
        :param pulumi.Input[str] private_key_password: The password to use in case the private key is encrypted.
        :param pulumi.Input[str] user: The user that we should use for the connection to the bastion host.
        """
        ProxyConnectionArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            host=host,
            agent_socket_path=agent_socket_path,
            dial_error_limit=dial_error_limit,
            password=password,
            per_dial_timeout=per_dial_timeout,
            port=port,
            private_key=private_key,
            private_key_password=private_key_password,
            user=user,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             host: Optional[pulumi.Input[str]] = None,
             agent_socket_path: Optional[pulumi.Input[str]] = None,
             dial_error_limit: Optional[pulumi.Input[int]] = None,
             password: Optional[pulumi.Input[str]] = None,
             per_dial_timeout: Optional[pulumi.Input[int]] = None,
             port: Optional[pulumi.Input[int]] = None,
             private_key: Optional[pulumi.Input[str]] = None,
             private_key_password: Optional[pulumi.Input[str]] = None,
             user: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if host is None:
            raise TypeError("Missing 'host' argument")
        if agent_socket_path is None and 'agentSocketPath' in kwargs:
            agent_socket_path = kwargs['agentSocketPath']
        if dial_error_limit is None and 'dialErrorLimit' in kwargs:
            dial_error_limit = kwargs['dialErrorLimit']
        if per_dial_timeout is None and 'perDialTimeout' in kwargs:
            per_dial_timeout = kwargs['perDialTimeout']
        if private_key is None and 'privateKey' in kwargs:
            private_key = kwargs['privateKey']
        if private_key_password is None and 'privateKeyPassword' in kwargs:
            private_key_password = kwargs['privateKeyPassword']

        _setter("host", host)
        if agent_socket_path is not None:
            _setter("agent_socket_path", agent_socket_path)
        if dial_error_limit is None:
            dial_error_limit = 10
        if dial_error_limit is not None:
            _setter("dial_error_limit", dial_error_limit)
        if password is not None:
            _setter("password", password)
        if per_dial_timeout is None:
            per_dial_timeout = 15
        if per_dial_timeout is not None:
            _setter("per_dial_timeout", per_dial_timeout)
        if port is None:
            port = 22
        if port is not None:
            _setter("port", port)
        if private_key is not None:
            _setter("private_key", private_key)
        if private_key_password is not None:
            _setter("private_key_password", private_key_password)
        if user is None:
            user = 'root'
        if user is not None:
            _setter("user", user)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Input[str]:
        """
        The address of the bastion host to connect to.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: pulumi.Input[str]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="agentSocketPath")
    def agent_socket_path(self) -> Optional[pulumi.Input[str]]:
        """
        SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
        """
        return pulumi.get(self, "agent_socket_path")

    @agent_socket_path.setter
    def agent_socket_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "agent_socket_path", value)

    @property
    @pulumi.getter(name="dialErrorLimit")
    def dial_error_limit(self) -> Optional[pulumi.Input[int]]:
        """
        Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
        """
        return pulumi.get(self, "dial_error_limit")

    @dial_error_limit.setter
    def dial_error_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dial_error_limit", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password we should use for the connection to the bastion host.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="perDialTimeout")
    def per_dial_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
        """
        return pulumi.get(self, "per_dial_timeout")

    @per_dial_timeout.setter
    def per_dial_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "per_dial_timeout", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port of the bastion host to connect to.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        The contents of an SSH key to use for the connection. This takes preference over the password if provided.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="privateKeyPassword")
    def private_key_password(self) -> Optional[pulumi.Input[str]]:
        """
        The password to use in case the private key is encrypted.
        """
        return pulumi.get(self, "private_key_password")

    @private_key_password.setter
    def private_key_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_password", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        The user that we should use for the connection to the bastion host.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


