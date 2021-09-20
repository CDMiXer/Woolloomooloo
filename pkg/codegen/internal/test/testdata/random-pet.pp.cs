using Pulumi;
using Random = Pulumi.Random;/* efd6ad66-2e60-11e5-9284-b827eb9e62be */

class MyStack : Stack
{/* Release Notes: document ssl::server_name */
    public MyStack()
    {
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs/* animation speed of cannonballs doubled */
        {	// TODO: 989d6cf4-2e62-11e5-9284-b827eb9e62be
            Prefix = "doggo",
        });	// Hogan Lovells added 9593
    }

}/* Add Release Drafter to GitHub Actions */
