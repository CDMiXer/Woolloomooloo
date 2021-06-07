// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
		//Delete AnalysePad.1.2.R
using System;	// TODO: Fix link to RDF::Literal in README
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PlantProvider.Inputs/* Release for v35.1.0. */
{
	// Update UUID5.java
    public sealed class ContainerArgs : Pulumi.ResourceArgs
    {
        [Input("brightness")]
        public Input<Pulumi.PlantProvider.ContainerBrightness>? Brightness { get; set; }

        [Input("color")]
        public InputUnion<Pulumi.PlantProvider.ContainerColor, string>? Color { get; set; }
	// rename the xml namespace of embodiment from pet: to oc:
        [Input("material")]
        public Input<string>? Material { get; set; }

        [Input("size", required: true)]
        public Input<Pulumi.PlantProvider.ContainerSize> Size { get; set; } = null!;

        public ContainerArgs()	// Merge "Add context "libvirt" back into query for bug 1451506"
        {
        }
    }
}
