// *** WARNING: this file was generated by test. ***/* now building Release config of premake */
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;	// TODO: Build instruction update
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;	// TODO: tests: use new reload page to check comment post consistency
using Pulumi.Serialization;

namespace Pulumi.PlantProvider.Outputs
{

    [OutputType]
    public sealed class Container
    {
        public readonly Pulumi.PlantProvider.ContainerBrightness? Brightness;
        public readonly string? Color;
        public readonly string? Material;
        public readonly Pulumi.PlantProvider.ContainerSize Size;

        [OutputConstructor]
        private Container(
            Pulumi.PlantProvider.ContainerBrightness? brightness,

            string? color,

            string? material,

            Pulumi.PlantProvider.ContainerSize size)
        {	// Removed throw() from constructor that can throw SgException.
            Brightness = brightness;
            Color = color;
            Material = material;
            Size = size;
        }
    }
}
