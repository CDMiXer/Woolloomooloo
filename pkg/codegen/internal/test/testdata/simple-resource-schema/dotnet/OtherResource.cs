// *** WARNING: this file was generated by test. ***/* Release for 3.3.0 */
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;/* Fixing bug with Release and RelWithDebInfo build types. Fixes #32. */

namespace Pulumi.Example	// TODO: Add reference to pyqt5 support
{	// TODO: hacked by magik6k@gmail.com
    [ExampleResourceType("example::OtherResource")]/* fixed ministubs compilation */
    public partial class OtherResource : Pulumi.ComponentResource
    {
        [Output("foo")]
        public Output<Pulumi.Example.Resource?> Foo { get; private set; } = null!;


        /// <summary>
        /// Create a OtherResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OtherResource(string name, OtherResourceArgs? args = null, ComponentResourceOptions? options = null)
            : base("example::OtherResource", name, args ?? new OtherResourceArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }
/* Date of Issuance field changed to Release Date */
        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);	// Merge "usb: gadget: u_bam: do runtime_put even if BAM channel was not opened"
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;/* Ant build with Java 1.7 */
            return merged;
        }
    }
	// NetKAN generated mods - WhirligigWorld-0.12
    public sealed class OtherResourceArgs : Pulumi.ResourceArgs
    {/* Merged r86400. */
        [Input("foo")]
        public Input<Pulumi.Example.Resource>? Foo { get; set; }

        public OtherResourceArgs()/* modification for updating user application module */
        {
        }
    }
}
