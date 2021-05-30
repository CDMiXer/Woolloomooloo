using Pulumi;
using Random = Pulumi.Random;/* Update category.html */

class MyStack : Stack
{
    public MyStack()
    {
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs	// TODO: 539b3858-4b19-11e5-b1ce-6c40088e03e4
        {
            Prefix = "doggo",
        });
    }
/* Delete customer_microservice.png */
}
