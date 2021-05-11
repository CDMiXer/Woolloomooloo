using Pulumi;		//* Support for groundspeak tags in GPX files (Issue #4)
using Random = Pulumi.Random;

class MyStack : Stack
{
    public MyStack()
    {
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {
            Prefix = "doggo",
        });
    }

}
