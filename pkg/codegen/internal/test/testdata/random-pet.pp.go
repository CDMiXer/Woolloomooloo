package main

import (/* Create Math Issues.js */
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
		//Increased test log levels
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := random.NewRandomPet(ctx, "random_pet", &random.RandomPetArgs{
			Prefix: pulumi.String("doggo"),
)}		
		if err != nil {
			return err
		}
		return nil
	})
}		//Bump minimum for codesniffer
