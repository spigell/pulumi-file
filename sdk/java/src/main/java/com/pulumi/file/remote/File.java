// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.file.remote;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.file.Utilities;
import com.pulumi.file.remote.FileArgs;
import com.pulumi.file.remote.outputs.Connection;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A file to be created on a remote host. The connection is established via ssh.
 * 
 */
@ResourceType(type="file:remote:File")
public class File extends com.pulumi.resources.CustomResource {
    /**
     * The parameters with which to connect to the remote host.
     * 
     */
    @Export(name="connection", refs={Connection.class}, tree="[0]")
    private Output<Connection> connection;

    /**
     * @return The parameters with which to connect to the remote host.
     * 
     */
    public Output<Connection> connection() {
        return this.connection;
    }
    /**
     * The content of file
     * 
     */
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> content;

    /**
     * @return The content of file
     * 
     */
    public Output<Optional<String>> content() {
        return Codegen.optional(this.content);
    }
    /**
     * The md5sum of the uploaded file
     * 
     */
    @Export(name="md5sum", refs={String.class}, tree="[0]")
    private Output<String> md5sum;

    /**
     * @return The md5sum of the uploaded file
     * 
     */
    public Output<String> md5sum() {
        return this.md5sum;
    }
    /**
     * The path for file on remote server
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
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
    @Export(name="permissions", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> permissions;

    /**
     * @return Unix permissions for file
     * 
     */
    public Output<Optional<String>> permissions() {
        return Codegen.optional(this.permissions);
    }
    /**
     * sudo mode requires a external sftp server to be running on remote host
     * 
     */
    @Export(name="sftpPath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sftpPath;

    /**
     * @return sudo mode requires a external sftp server to be running on remote host
     * 
     */
    public Output<Optional<String>> sftpPath() {
        return Codegen.optional(this.sftpPath);
    }
    /**
     * Trigger replacements on changes to this input.
     * 
     */
    @Export(name="triggers", refs={List.class,Object.class}, tree="[0,1]")
    private Output</* @Nullable */ List<Object>> triggers;

    /**
     * @return Trigger replacements on changes to this input.
     * 
     */
    public Output<Optional<List<Object>>> triggers() {
        return Codegen.optional(this.triggers);
    }
    /**
     * if enabled then use sudo for copy command instead of direct copy
     * 
     */
    @Export(name="useSudo", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> useSudo;

    /**
     * @return if enabled then use sudo for copy command instead of direct copy
     * 
     */
    public Output<Optional<Boolean>> useSudo() {
        return Codegen.optional(this.useSudo);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public File(String name) {
        this(name, FileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public File(String name, FileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public File(String name, FileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("file:remote:File", name, args == null ? FileArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private File(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("file:remote:File", name, null, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "connection"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static File get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new File(name, id, options);
    }
}
