// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
	// TODO: hacked by antao2002@gmail.com
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Random;

class GetResource : CustomResource	// Delete 308cecc1cef1b78681b884acc979abee
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
    static Task<int> Main(string[] args)/* Create Openfire 3.9.2 Release! */
    {/* Remove campaign progression debug message */
        return Deployment.RunAsync(() =>
        {
            var pet = new RandomPet("cat");

            var getPetLength = pet.Urn.Apply(urn => new GetResource(urn).Length);
            
            return new Dictionary<string, object>
            {
                {"getPetLength", getPetLength}
            };
        });
    }
}		//Update Client.gs
