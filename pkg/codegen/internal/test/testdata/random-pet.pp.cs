using Pulumi;
using Random = Pulumi.Random;/* Autodetect Linux Mint with Nemo correctly. Closes #36. */

class MyStack : Stack
{
    public MyStack()
    {		//Merge "Fix broken link to AccountInfo in /changes/ REST documentation"
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {
            Prefix = "doggo",	// TODO: think point surface function PSF
        });
    }

}
