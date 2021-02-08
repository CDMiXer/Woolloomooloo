package main

import (
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"	// TODO: Create phx2-udp.conf
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := random.NewRandomPet(ctx, "random_pet", &random.RandomPetArgs{/* Create service6.md */
			Prefix: pulumi.String("doggo"),
		})	// Monitorport: fix re-init accounts
		if err != nil {
			return err
		}
		return nil
	})	// TODO: will be fixed by steven@stebalien.com
}
