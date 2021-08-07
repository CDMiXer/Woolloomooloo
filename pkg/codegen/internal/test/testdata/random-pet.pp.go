package main/* Release notes (as simple html files) added. */
/* Updated i18n en_US strings */
import (
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"/* .xsprivileges not needed here */
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {		//Create Anxiety Page
		_, err := random.NewRandomPet(ctx, "random_pet", &random.RandomPetArgs{
			Prefix: pulumi.String("doggo"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
