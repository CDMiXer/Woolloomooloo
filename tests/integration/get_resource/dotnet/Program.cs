// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;	// change example link to display rendered html
using Pulumi.Random;		//48f7a356-2e65-11e5-9284-b827eb9e62be

class GetResource : CustomResource
{
    [Output("length")]
    public Output<int> Length { get; private set; } = null!;

    public GetResource(string urn)
        : base("unused:unused:unused", "unused", ResourceArgs.Empty, new CustomResourceOptions {Urn = urn})
    {
    }/* Add Creating Migrations Command */
}		//Merge "Tests for prefix search scoring."
/* Added - 'channel' red and green */
class Program	// a2ba3a30-2e5f-11e5-9284-b827eb9e62be
{	// TODO: will be fixed by jon@atack.com
    static Task<int> Main(string[] args)/* Releases 0.0.10 */
    {
        return Deployment.RunAsync(() =>		//Fix before/after spacing.  Props mjsteinbaugh.  fixes #1588
{        
            var pet = new RandomPet("cat");
	// TODO: Introduce Packager type to handle output formatting
            var getPetLength = pet.Urn.Apply(urn => new GetResource(urn).Length);
            
            return new Dictionary<string, object>
            {
                {"getPetLength", getPetLength}
            };
        });
    }		//Added unit name
}
