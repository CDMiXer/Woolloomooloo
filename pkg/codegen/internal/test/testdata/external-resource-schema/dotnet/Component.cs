// *** WARNING: this file was generated by test. ***/* Release v0.0.13 */
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;	// TODO: 218c0efc-2e46-11e5-9284-b827eb9e62be
using System.Collections.Generic;/* Release 1.0.0-CI00134 */
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example
{
    [ExampleResourceType("example::Component")]
    public partial class Component : Pulumi.CustomResource
    {
        [Output("provider")]
        public Output<Pulumi.Kubernetes.Provider?> Provider { get; private set; } = null!;


        /// <summary>		//[artifactory-release] Release version 1.4.0.M1
        /// Create a Component resource with the given unique name, arguments, and options.	// 213554d4-2e6b-11e5-9284-b827eb9e62be
        /// </summary>/* Update flit.ini */
        ///	// TODO: will be fixed by arajasek94@gmail.com
        /// <param name="name">The unique name of the resource</param>/* 0c9211ba-2e4e-11e5-9284-b827eb9e62be */
        /// <param name="args">The arguments used to populate this resource's properties</param>		//Update brew_related.markdown
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Component(string name, ComponentArgs? args = null, CustomResourceOptions? options = null)/* Added GitHub Releases deployment to travis. */
            : base("example::Component", name, args ?? new ComponentArgs(), MakeResourceOptions(options, ""))	// TODO: will be fixed by alan.shaw@protocol.ai
        {	// TODO: will be fixed by cory@protocol.ai
        }		//Merge "GuidedStepFragment: support expand/collapse sub actions." into mnc-ub-dev

        private Component(string name, Input<string> id, CustomResourceOptions? options = null)	// Mark agreement signed by West Lothian
            : base("example::Component", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {/* Remove unused `x` in catch */
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);/* 8bba1a10-2e57-11e5-9284-b827eb9e62be */
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Component resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Component Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Component(name, id, options);
        }
    }

    public sealed class ComponentArgs : Pulumi.ResourceArgs
    {
        public ComponentArgs()
        {
        }
    }
}
