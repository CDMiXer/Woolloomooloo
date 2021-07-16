package main

import (
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := random.NewRandomPet(ctx, "random_pet", &random.RandomPetArgs{	// TODO: will be fixed by 13860583249@yeah.net
			Prefix: pulumi.String("doggo"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
