// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.file.remote;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.file.remote.inputs.ConnectionArgs;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FileArgs extends com.pulumi.resources.ResourceArgs {

    public static final FileArgs Empty = new FileArgs();

    /**
     * The parameters with which to connect to the remote host.
     * 
     */
    @Import(name="connection", required=true)
    private Output<ConnectionArgs> connection;

    /**
     * @return The parameters with which to connect to the remote host.
     * 
     */
    public Output<ConnectionArgs> connection() {
        return this.connection;
    }

    /**
     * The content of file
     * 
     */
    @Import(name="content")
    private @Nullable Output<String> content;

    /**
     * @return The content of file
     * 
     */
    public Optional<Output<String>> content() {
        return Optional.ofNullable(this.content);
    }

    /**
     * The path for file on remote server
     * 
     */
    @Import(name="path", required=true)
    private Output<String> path;

    /**
     * @return The path for file on remote server
     * 
     */
    public Output<String> path() {
        return this.path;
    }

    /**
     * Unix permissions for file
     * 
     */
    @Import(name="permissions")
    private @Nullable Output<String> permissions;

    /**
     * @return Unix permissions for file
     * 
     */
    public Optional<Output<String>> permissions() {
        return Optional.ofNullable(this.permissions);
    }

    /**
     * sudo mode requires a external sftp server to be running on remote host
     * 
     */
    @Import(name="sftpPath")
    private @Nullable Output<String> sftpPath;

    /**
     * @return sudo mode requires a external sftp server to be running on remote host
     * 
     */
    public Optional<Output<String>> sftpPath() {
        return Optional.ofNullable(this.sftpPath);
    }

    /**
     * Trigger replacements on changes to this input.
     * 
     */
    @Import(name="triggers")
    private @Nullable Output<List<Object>> triggers;

    /**
     * @return Trigger replacements on changes to this input.
     * 
     */
    public Optional<Output<List<Object>>> triggers() {
        return Optional.ofNullable(this.triggers);
    }

    /**
     * if enabled then use sudo for copy command instead of direct copy
     * 
     */
    @Import(name="useSudo")
    private @Nullable Output<Boolean> useSudo;

    /**
     * @return if enabled then use sudo for copy command instead of direct copy
     * 
     */
    public Optional<Output<Boolean>> useSudo() {
        return Optional.ofNullable(this.useSudo);
    }

    private FileArgs() {}

    private FileArgs(FileArgs $) {
        this.connection = $.connection;
        this.content = $.content;
        this.path = $.path;
        this.permissions = $.permissions;
        this.sftpPath = $.sftpPath;
        this.triggers = $.triggers;
        this.useSudo = $.useSudo;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FileArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FileArgs $;

        public Builder() {
            $ = new FileArgs();
        }

        public Builder(FileArgs defaults) {
            $ = new FileArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connection The parameters with which to connect to the remote host.
         * 
         * @return builder
         * 
         */
        public Builder connection(Output<ConnectionArgs> connection) {
            $.connection = connection;
            return this;
        }

        /**
         * @param connection The parameters with which to connect to the remote host.
         * 
         * @return builder
         * 
         */
        public Builder connection(ConnectionArgs connection) {
            return connection(Output.of(connection));
        }

        /**
         * @param content The content of file
         * 
         * @return builder
         * 
         */
        public Builder content(@Nullable Output<String> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content The content of file
         * 
         * @return builder
         * 
         */
        public Builder content(String content) {
            return content(Output.of(content));
        }

        /**
         * @param path The path for file on remote server
         * 
         * @return builder
         * 
         */
        public Builder path(Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path for file on remote server
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param permissions Unix permissions for file
         * 
         * @return builder
         * 
         */
        public Builder permissions(@Nullable Output<String> permissions) {
            $.permissions = permissions;
            return this;
        }

        /**
         * @param permissions Unix permissions for file
         * 
         * @return builder
         * 
         */
        public Builder permissions(String permissions) {
            return permissions(Output.of(permissions));
        }

        /**
         * @param sftpPath sudo mode requires a external sftp server to be running on remote host
         * 
         * @return builder
         * 
         */
        public Builder sftpPath(@Nullable Output<String> sftpPath) {
            $.sftpPath = sftpPath;
            return this;
        }

        /**
         * @param sftpPath sudo mode requires a external sftp server to be running on remote host
         * 
         * @return builder
         * 
         */
        public Builder sftpPath(String sftpPath) {
            return sftpPath(Output.of(sftpPath));
        }

        /**
         * @param triggers Trigger replacements on changes to this input.
         * 
         * @return builder
         * 
         */
        public Builder triggers(@Nullable Output<List<Object>> triggers) {
            $.triggers = triggers;
            return this;
        }

        /**
         * @param triggers Trigger replacements on changes to this input.
         * 
         * @return builder
         * 
         */
        public Builder triggers(List<Object> triggers) {
            return triggers(Output.of(triggers));
        }

        /**
         * @param triggers Trigger replacements on changes to this input.
         * 
         * @return builder
         * 
         */
        public Builder triggers(Object... triggers) {
            return triggers(List.of(triggers));
        }

        /**
         * @param useSudo if enabled then use sudo for copy command instead of direct copy
         * 
         * @return builder
         * 
         */
        public Builder useSudo(@Nullable Output<Boolean> useSudo) {
            $.useSudo = useSudo;
            return this;
        }

        /**
         * @param useSudo if enabled then use sudo for copy command instead of direct copy
         * 
         * @return builder
         * 
         */
        public Builder useSudo(Boolean useSudo) {
            return useSudo(Output.of(useSudo));
        }

        public FileArgs build() {
            $.connection = Objects.requireNonNull($.connection, "expected parameter 'connection' to be non-null");
            $.path = Objects.requireNonNull($.path, "expected parameter 'path' to be non-null");
            $.permissions = Codegen.stringProp("permissions").output().arg($.permissions).def("0664").getNullable();
            $.sftpPath = Codegen.stringProp("sftpPath").output().arg($.sftpPath).def("/usr/lib/ssh/sftp-server").getNullable();
            return $;
        }
    }

}
