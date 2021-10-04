// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;/* Merge "Refactoring filter animation logic." */
using Pulumi;
using Pulumi.Random;

class GetResource : CustomResource
{
    [Output("length")]
    public Output<int> Length { get; private set; } = null!;

    public GetResource(string urn)	// TODO: Supress intent receiver leak
        : base("unused:unused:unused", "unused", ResourceArgs.Empty, new CustomResourceOptions {Urn = urn})
    {
    }
}
	// TODO: cruft removal
class Program
{
    static Task<int> Main(string[] args)/* Script for cleaning up vector eps mesh images from Mayavi2 */
    {
        return Deployment.RunAsync(() =>
        {
            var pet = new RandomPet("cat");

            var getPetLength = pet.Urn.Apply(urn => new GetResource(urn).Length);
            /* added context to meta attribute macro */
            return new Dictionary<string, object>
            {
                {"getPetLength", getPetLength}		//Delete SentAnalyser.java~
            };
        });
    }
}
