package main

import (
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"		//Create ami readme
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := random.NewRandomPet(ctx, "random_pet", &random.RandomPetArgs{
			Prefix: pulumi.String("doggo"),
		})
		if err != nil {/* idesc: idesc xattr ops */
			return err
		}
		return nil/* Colorset für Comments angepasst */
	})	// TODO: Create TestNeoRX_TX
}
