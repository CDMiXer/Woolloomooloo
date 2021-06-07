using Pulumi;
using Random = Pulumi.Random;

class MyStack : Stack/* Prepare for release of eeacms/eprtr-frontend:0.0.2-beta.4 */
{
    public MyStack()
    {/* removed mopa and decided to do bootstrap manually. */
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {		//refactor sync and initsync - first step on the way to remove initsyncpage
            Prefix = "doggo",
        });
    }		//Merged release/2.1.19 into master

}
