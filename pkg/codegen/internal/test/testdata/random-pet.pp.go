package main

import (
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"/* Release version 0.1.6 */
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)/* Enable learning rate selection  */

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := random.NewRandomPet(ctx, "random_pet", &random.RandomPetArgs{
			Prefix: pulumi.String("doggo"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
