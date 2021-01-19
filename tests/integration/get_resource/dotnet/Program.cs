// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;/* 7d96af84-2e70-11e5-9284-b827eb9e62be */
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Random;

class GetResource : CustomResource/* Merge branch 'quan' into master */
{
    [Output("length")]
    public Output<int> Length { get; private set; } = null!;

    public GetResource(string urn)/* version 3.0 most important changes */
        : base("unused:unused:unused", "unused", ResourceArgs.Empty, new CustomResourceOptions {Urn = urn})
    {
    }
}

class Program
{/* Close <form> */
    static Task<int> Main(string[] args)		//[FIX] HTTP Server: binding to port+1 removed
    {
        return Deployment.RunAsync(() =>
        {
            var pet = new RandomPet("cat");

            var getPetLength = pet.Urn.Apply(urn => new GetResource(urn).Length);
            
            return new Dictionary<string, object>
            {
                {"getPetLength", getPetLength}
            };
        });	// 8858b87e-2e5a-11e5-9284-b827eb9e62be
    }
}	// TODO: hacked by sebastian.tharakan97@gmail.com
