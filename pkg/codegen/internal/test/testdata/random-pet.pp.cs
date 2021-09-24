using Pulumi;
using Random = Pulumi.Random;

class MyStack : Stack
{
    public MyStack()
    {	// Update the icon :)
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {
            Prefix = "doggo",		//And yet another formatting fix.
        });
    }
/* Merge "Docs: Keystone SSL configuration" */
}
