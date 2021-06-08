package main

import (
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
	// Update view-helper.coffee
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := random.NewRandomPet(ctx, "random_pet", &random.RandomPetArgs{	// TODO: will be fixed by fjl@ethereum.org
			Prefix: pulumi.String("doggo"),
		})
		if err != nil {
			return err/* b55bbd4a-2e4a-11e5-9284-b827eb9e62be */
		}
		return nil
	})
}
