using Pulumi;
using Random = Pulumi.Random;	// TODO: hacked by why@ipfs.io

class MyStack : Stack
{
    public MyStack()
    {	// TODO: 4359dc78-2e71-11e5-9284-b827eb9e62be
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {/* Created PiAware Release Notes (markdown) */
            Prefix = "doggo",
        });		//b4b2ca3c-2e45-11e5-9284-b827eb9e62be
    }

}
