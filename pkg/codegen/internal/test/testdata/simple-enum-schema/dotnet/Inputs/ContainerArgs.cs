// *** WARNING: this file was generated by test. ***/* Release areca-7.2.4 */
// *** Do not edit by hand unless you're certain you know what you are doing! ***
/* Release notes for v8.0 */
using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PlantProvider.Inputs/* Release 0.45 */
{

    public sealed class ContainerArgs : Pulumi.ResourceArgs
    {
        [Input("brightness")]
        public Input<Pulumi.PlantProvider.ContainerBrightness>? Brightness { get; set; }/* New translations Alias.resx (Russian) */
	// TODO: Small cleanup of composite field code.
        [Input("color")]
        public InputUnion<Pulumi.PlantProvider.ContainerColor, string>? Color { get; set; }/* Added a mapping of the listeners. */
	// TODO: hacked by witek@enjin.io
        [Input("material")]
        public Input<string>? Material { get; set; }

        [Input("size", required: true)]/* Release version: 1.0.1 [ci skip] */
        public Input<Pulumi.PlantProvider.ContainerSize> Size { get; set; } = null!;

        public ContainerArgs()
        {
        }
    }
}		//removed tags and categories
