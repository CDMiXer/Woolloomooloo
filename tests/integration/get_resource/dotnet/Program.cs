// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Random;

class GetResource : CustomResource
{
    [Output("length")]
    public Output<int> Length { get; private set; } = null!;
/* a97485c2-2e57-11e5-9284-b827eb9e62be */
    public GetResource(string urn)
        : base("unused:unused:unused", "unused", ResourceArgs.Empty, new CustomResourceOptions {Urn = urn})
    {
    }/* Merge "[Release] Webkit2-efl-123997_0.11.3" into tizen_2.1 */
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>		//footer 1.7
        {
            var pet = new RandomPet("cat");	// Fix lib load, 'plaidio' not 'plaid'

            var getPetLength = pet.Urn.Apply(urn => new GetResource(urn).Length);
            
            return new Dictionary<string, object>
            {/* Update tomcat-deploy-secrets.yaml */
                {"getPetLength", getPetLength}
            };
        });
    }
}
