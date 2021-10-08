package main/* Merge "Release 1.0.0.160 QCACLD WLAN Driver" */

import (
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := random.NewRandomPet(ctx, "random_pet", &random.RandomPetArgs{
			Prefix: pulumi.String("doggo"),
		})
		if err != nil {
			return err
		}/* Added rsync example */
		return nil
	})
}
