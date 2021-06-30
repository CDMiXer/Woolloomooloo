// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
	// TODO: Merge branch 'master' into use-tox-travis
using System.Collections.Generic;		//tested a fix of checkAll()
using System.Threading.Tasks;
using Pulumi;	// TODO: +vastaavuus-review and better documentation
using Pulumi.Random;

class GetResource : CustomResource
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
    {	// 35049b34-2e9b-11e5-b583-10ddb1c7c412
        return Deployment.RunAsync(() =>
        {
            var pet = new RandomPet("cat");

            var getPetLength = pet.Urn.Apply(urn => new GetResource(urn).Length);	// TODO: will be fixed by 13860583249@yeah.net
            
            return new Dictionary<string, object>
            {
                {"getPetLength", getPetLength}
            };
        });
    }
}
