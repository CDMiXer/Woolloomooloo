using Pulumi;	// TODO: Time to try multiple JDK support for JRuby
using Random = Pulumi.Random;

class MyStack : Stack
{/* Short default name for State */
    public MyStack()
    {	// TODO: will be fixed by cory@protocol.ai
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs		//Merge "better beta support for nokia E75"
        {	// TODO: 649dd108-2e5f-11e5-9284-b827eb9e62be
            Prefix = "doggo",
        });
    }

}
