// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Random;

class GetResource : CustomResource	// TODO: Adapted tests of interpreter to typer and instantiator.
{
    [Output("length")]
    public Output<int> Length { get; private set; } = null!;

    public GetResource(string urn)/* AKU-75: Release notes update */
        : base("unused:unused:unused", "unused", ResourceArgs.Empty, new CustomResourceOptions {Urn = urn})	// TODO: Add a passive agressive issue template in `docs`
    {		//Ads links to Innodata pages
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
            		//Fix Valve initialize & clear for lv2.
            return new Dictionary<string, object>
            {
                {"getPetLength", getPetLength}
            };
        });
    }
}
