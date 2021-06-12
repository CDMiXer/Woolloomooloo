using Pulumi;/* Add symbol as true to read me	 */
using Random = Pulumi.Random;	// 1f4c2f5e-2d5c-11e5-8442-b88d120fff5e

class MyStack : Stack
{
    public MyStack()
    {
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {
            Prefix = "doggo",
        });
    }

}	// TODO: hacked by why@ipfs.io
