// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.File.Remote
{
    /// <summary>
    /// A file to be created on a remote host. The connection is established via ssh.
    /// </summary>
    [FileResourceType("file:remote:File")]
    public partial class File : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The parameters with which to connect to the remote host.
        /// </summary>
        [Output("connection")]
        public Output<Outputs.Connection> Connection { get; private set; } = null!;

        /// <summary>
        /// The content of file
        /// </summary>
        [Output("content")]
        public Output<string?> Content { get; private set; } = null!;

        /// <summary>
        /// The md5sum of the uploaded file
        /// </summary>
        [Output("md5sum")]
        public Output<string> Md5sum { get; private set; } = null!;

        /// <summary>
        /// The path for file on remote server
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Unix permissions for file
        /// </summary>
        [Output("permissions")]
        public Output<string?> Permissions { get; private set; } = null!;

        /// <summary>
        /// sudo mode requires a external sftp server to be running on remote host
        /// </summary>
        [Output("sftpPath")]
        public Output<string?> SftpPath { get; private set; } = null!;

        /// <summary>
        /// Trigger replacements on changes to this input.
        /// </summary>
        [Output("triggers")]
        public Output<ImmutableArray<object>> Triggers { get; private set; } = null!;

        /// <summary>
        /// if enabled then use sudo for copy command instead of direct copy
        /// </summary>
        [Output("useSudo")]
        public Output<bool?> UseSudo { get; private set; } = null!;


        /// <summary>
        /// Create a File resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public File(string name, FileArgs args, CustomResourceOptions? options = null)
            : base("file:remote:File", name, args ?? new FileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private File(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("file:remote:File", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/spigell/pulumi-file",
                AdditionalSecretOutputs =
                {
                    "connection",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing File resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static File Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new File(name, id, options);
        }
    }

    public sealed class FileArgs : global::Pulumi.ResourceArgs
    {
        [Input("connection", required: true)]
        private Input<Inputs.ConnectionArgs>? _connection;

        /// <summary>
        /// The parameters with which to connect to the remote host.
        /// </summary>
        public Input<Inputs.ConnectionArgs>? Connection
        {
            get => _connection;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _connection = Output.Tuple<Input<Inputs.ConnectionArgs>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The content of file
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The path for file on remote server
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Unix permissions for file
        /// </summary>
        [Input("permissions")]
        public Input<string>? Permissions { get; set; }

        /// <summary>
        /// sudo mode requires a external sftp server to be running on remote host
        /// </summary>
        [Input("sftpPath")]
        public Input<string>? SftpPath { get; set; }

        [Input("triggers")]
        private InputList<object>? _triggers;

        /// <summary>
        /// Trigger replacements on changes to this input.
        /// </summary>
        public InputList<object> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<object>());
            set => _triggers = value;
        }

        /// <summary>
        /// if enabled then use sudo for copy command instead of direct copy
        /// </summary>
        [Input("useSudo")]
        public Input<bool>? UseSudo { get; set; }

        public FileArgs()
        {
            Permissions = "0664";
            SftpPath = "/usr/lib/ssh/sftp-server";
        }
        public static new FileArgs Empty => new FileArgs();
    }
}
