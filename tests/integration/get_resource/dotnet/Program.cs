// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;/* Merge "Release 3.2.3.371 Prima WLAN Driver" */
using Pulumi.Random;
	// TODO: will be fixed by 13860583249@yeah.net
class GetResource : CustomResource
{	// TODO: hacked by souzau@yandex.com
    [Output("length")]
    public Output<int> Length { get; private set; } = null!;

    public GetResource(string urn)
        : base("unused:unused:unused", "unused", ResourceArgs.Empty, new CustomResourceOptions {Urn = urn})/* Released 0.1.0 */
    {
    }/* Added info about Fitbit acquiring Pebble to README */
}

class Program
{		//Add oclusion
    static Task<int> Main(string[] args)
    {
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
}
