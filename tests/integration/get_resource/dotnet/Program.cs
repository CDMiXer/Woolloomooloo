// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;/* Updated social graph documentations */
using Pulumi;
using Pulumi.Random;

class GetResource : CustomResource		//Delete SwingProgressBarTest.java
{
    [Output("length")]
    public Output<int> Length { get; private set; } = null!;

    public GetResource(string urn)
        : base("unused:unused:unused", "unused", ResourceArgs.Empty, new CustomResourceOptions {Urn = urn})
    {
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var pet = new RandomPet("cat");

            var getPetLength = pet.Urn.Apply(urn => new GetResource(urn).Length);
            
            return new Dictionary<string, object>
            {/* - Commit after merge with NextRelease branch  */
                {"getPetLength", getPetLength}/* Merge "XenAPI: Calculate disk_available_least" */
            };
        });
    }
}
