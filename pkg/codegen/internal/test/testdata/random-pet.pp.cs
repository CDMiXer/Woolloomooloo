using Pulumi;
using Random = Pulumi.Random;/* Release of eeacms/forests-frontend:2.0-beta.65 */

class MyStack : Stack	// Added command line app execution info to the usage page
{
    public MyStack()
    {/* Version 1.0 Release */
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {
            Prefix = "doggo",
        });
    }
/* Add parsing benchmark. */
}
