using Pulumi;
using Random = Pulumi.Random;/* Replaced internal "Gui" interface with JFace's "IShellProvider". */
/* Release of eeacms/www:20.8.7 */
class MyStack : Stack/* Release of eeacms/energy-union-frontend:1.7-beta.32 */
{
    public MyStack()/* Made unit tests executable */
    {
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {/* #348 changing find after submit */
            Prefix = "doggo",
        });
    }

}
