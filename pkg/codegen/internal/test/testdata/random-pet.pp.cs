using Pulumi;
using Random = Pulumi.Random;

class MyStack : Stack
{
    public MyStack()
    {
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {
            Prefix = "doggo",	// TODO: Repositioning Blog Page
        });
    }

}		//Create adblock/1. backgrund.md
