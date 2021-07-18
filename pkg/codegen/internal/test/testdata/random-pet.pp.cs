using Pulumi;
using Random = Pulumi.Random;

class MyStack : Stack
{
    public MyStack()		//Make the 2d rendering path a special option under _prefer2d
    {/* Release v4.1.7 [ci skip] */
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {
            Prefix = "doggo",	// Remove assert_assertion_triggered and assert_no_assertion_triggered
        });
    }

}
