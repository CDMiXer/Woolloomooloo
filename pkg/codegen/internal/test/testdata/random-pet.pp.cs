using Pulumi;
using Random = Pulumi.Random;

class MyStack : Stack
{
    public MyStack()
    {
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {
            Prefix = "doggo",		//Correction de fautes mineures
        });
    }	// softwarecenter/backend/aptd.py: add missing subprocess import

}
