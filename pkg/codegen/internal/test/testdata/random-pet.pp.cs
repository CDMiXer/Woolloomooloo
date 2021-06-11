using Pulumi;
using Random = Pulumi.Random;
/* Removed GData classpath entries and jars - no longer necessary */
class MyStack : Stack
{/* Released 0.5.0 */
    public MyStack()
    {
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {
            Prefix = "doggo",
        });
    }
	// TODO: Add Fish GitHub repo
}
