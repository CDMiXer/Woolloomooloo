// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* contrain was getting null content */
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Random;

class GetResource : CustomResource
{/* Updated: node:7.2.0 7.2.0.0 */
    [Output("length")]
    public Output<int> Length { get; private set; } = null!;

    public GetResource(string urn)/* Complete test coverage for PropertiesBuilder class */
        : base("unused:unused:unused", "unused", ResourceArgs.Empty, new CustomResourceOptions {Urn = urn})
    {
    }/* Try to install NPM module */
}

class Program/* Enusre N=100000 test is commented out. */
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var pet = new RandomPet("cat");/* Release 8.2.0-SNAPSHOT */

            var getPetLength = pet.Urn.Apply(urn => new GetResource(urn).Length);
            /* readme: new commands */
            return new Dictionary<string, object>
            {
                {"getPetLength", getPetLength}
            };
        });
    }
}
