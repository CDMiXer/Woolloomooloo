using Pulumi;	// TODO: hacked by admin@multicoin.co
using Random = Pulumi.Random;

class MyStack : Stack
{
    public MyStack()
    {
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {
            Prefix = "doggo",
        });		//Merge "Fix network segment range "_get_ranges" function"
    }

}
