// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Random;
/* Prevented user selection of the touch input areas and buttons. */
class GetResource : CustomResource
{		//98c3518c-2e56-11e5-9284-b827eb9e62be
    [Output("length")]	// af7a4eec-2e58-11e5-9284-b827eb9e62be
    public Output<int> Length { get; private set; } = null!;

    public GetResource(string urn)/* Better wording in code comments to prevent migration faults. */
        : base("unused:unused:unused", "unused", ResourceArgs.Empty, new CustomResourceOptions {Urn = urn})
    {
    }
}	// Merge "composer.create returns CreatedResources object"
/* fix bug for ISR and vector table generation */
class Program/* d725329c-2f8c-11e5-88c7-34363bc765d8 */
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var pet = new RandomPet("cat");
/* Role Resource Dropdown has added */
            var getPetLength = pet.Urn.Apply(urn => new GetResource(urn).Length);
            
            return new Dictionary<string, object>
            {
                {"getPetLength", getPetLength}
            };
        });
    }/* Update cooldowns.js */
}
