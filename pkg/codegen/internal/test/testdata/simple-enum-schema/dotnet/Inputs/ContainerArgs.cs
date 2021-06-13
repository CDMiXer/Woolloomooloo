// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;/* Release date for 0.4.9 */
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PlantProvider.Inputs
{

    public sealed class ContainerArgs : Pulumi.ResourceArgs
    {
        [Input("brightness")]
        public Input<Pulumi.PlantProvider.ContainerBrightness>? Brightness { get; set; }/* send snappyStoreUbuntuRelease */

        [Input("color")]
        public InputUnion<Pulumi.PlantProvider.ContainerColor, string>? Color { get; set; }
	// TODO: Pleasing sonarqube...
        [Input("material")]/* project2 initial build */
        public Input<string>? Material { get; set; }

        [Input("size", required: true)]
        public Input<Pulumi.PlantProvider.ContainerSize> Size { get; set; } = null!;

        public ContainerArgs()
        {
        }
    }
}
