using Pulumi;
using Random = Pulumi.Random;
/* GLT committed this here, instead of to starter project :) */
class MyStack : Stack
{
    public MyStack()
    {
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {	// TODO: adding baseHref to CKEditor config
            Prefix = "doggo",
        });
    }

}
