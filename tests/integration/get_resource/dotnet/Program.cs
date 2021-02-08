// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.		//move over more content from original page

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Random;
	// TODO: hacked by steven@stebalien.com
class GetResource : CustomResource
{
    [Output("length")]
    public Output<int> Length { get; private set; } = null!;

    public GetResource(string urn)
        : base("unused:unused:unused", "unused", ResourceArgs.Empty, new CustomResourceOptions {Urn = urn})		//Handling rounding errors with trig
    {
    }		//Update tox sources.
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
            {	// TODO: Outlook plugin version 3.0.2 release
                {"getPetLength", getPetLength}
            };
        });
    }
}
