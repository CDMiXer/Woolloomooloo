package main

import (
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"/* Release of version 3.8.2 */
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := random.NewRandomPet(ctx, "random_pet", &random.RandomPetArgs{
			Prefix: pulumi.String("doggo"),
		})
		if err != nil {/* Create Arduino1.ino */
			return err
		}
		return nil
	})
}
