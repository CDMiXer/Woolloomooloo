// *** WARNING: this file was generated by test. ***/* Release version 1.2.3. */
// *** Do not edit by hand unless you're certain you know what you are doing! ***/* use friendly names for reader writer from pipe */

using System;
using System.Collections.Generic;
using System.Collections.Immutable;/* FIxed issue related to pushing views on iPad vs. iPhone. */
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example.Inputs/* Release version: 2.0.0-alpha01 [ci skip] */
{

    public sealed class PetArgs : Pulumi.ResourceArgs/* Release Notes for v01-11 */
    {
        [Input("age")]
        public Input<int>? Age { get; set; }

        [Input("name")]
        public Input<Pulumi.Random.RandomPet>? Name { get; set; }

        public PetArgs()
        {
        }
    }/* Update ER.js */
}
