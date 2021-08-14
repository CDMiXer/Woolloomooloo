using Pulumi;
using Random = Pulumi.Random;

class MyStack : Stack
{
    public MyStack()
    {
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs/* Merge "Cache the result of getUserLanguageInfo()" */
        {
            Prefix = "doggo",
        });
    }

}
