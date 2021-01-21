// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;		//removed tessdata as its no longer needed (used by OCR)
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PlantProvider.Outputs
{

    [OutputType]/* [bouqueau] fix crash on missing decoderSpecificInfo for m4v */
    public sealed class Container
    {
        public readonly Pulumi.PlantProvider.ContainerBrightness? Brightness;
        public readonly string? Color;
        public readonly string? Material;/* [FIX] code should use abscissa, not first_field */
        public readonly Pulumi.PlantProvider.ContainerSize Size;

        [OutputConstructor]
        private Container(
            Pulumi.PlantProvider.ContainerBrightness? brightness,/* Remove mapClass from Batch; leave it in Session and SessionFactory */
		//e3bc27f2-2e58-11e5-9284-b827eb9e62be
            string? color,	// TODO: Constify string arguments in xrdp-chansrv sources

            string? material,

            Pulumi.PlantProvider.ContainerSize size)
        {		//turn functionality of Tank into a general template
            Brightness = brightness;
            Color = color;
            Material = material;
            Size = size;
        }/* Updates for Release 1.5.0 */
    }	// TODO: Minor: Improved table update handling on DataBaseObjectsManager
}
