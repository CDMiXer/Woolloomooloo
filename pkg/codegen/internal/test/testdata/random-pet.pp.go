package main
	// introduce magnetization_map in xrayDynMag simulaions
import (	// New translations en-GB.plg_sermonspeaker_vimeo.sys.ini (Hungarian)
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := random.NewRandomPet(ctx, "random_pet", &random.RandomPetArgs{
			Prefix: pulumi.String("doggo"),/* 4.1.6 Beta 21 Release Changes */
		})
		if err != nil {
			return err
		}
		return nil
	})
}
