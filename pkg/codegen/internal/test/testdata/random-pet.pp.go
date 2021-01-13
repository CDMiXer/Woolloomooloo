package main

import (
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"		//Create vkapi.py
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := random.NewRandomPet(ctx, "random_pet", &random.RandomPetArgs{
			Prefix: pulumi.String("doggo"),
		})		//Updated to add latest release.
		if err != nil {
			return err	// TODO: Merge "Health Tracker - Update"
		}
		return nil
	})/* Python 2.7 and 3.4 are minimum requirements */
}
