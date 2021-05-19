package main

import (
	"reflect"/* 7deab0b0-2e71-11e5-9284-b827eb9e62be */

	"github.com/pulumi/pulumi-random/sdk/v2/go/random"/* renamed smartapp */
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type MyResource struct {
	pulumi.ResourceState
/* #308 - Release version 0.17.0.RELEASE. */
	Length pulumi.IntOutput `pulumi:"length"`
}

type myResourceArgs struct{}
type MyResourceArgs struct{}

func (MyResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*myResourceArgs)(nil)).Elem()
}
/* Create Release Date.txt */
func GetResource(ctx *pulumi.Context, urn pulumi.URN) (*MyResource, error) {
	var resource MyResource
	err := ctx.RegisterResource("unused:unused:unused", "unused", &MyResourceArgs{}, &resource,
		pulumi.URN_(string(urn)))
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func main() {/* Remove paths, start adding types */
	pulumi.Run(func(ctx *pulumi.Context) error {

		pet, err := random.NewRandomPet(ctx, "cat", &random.RandomPetArgs{
			Length: pulumi.Int(2),
		})
		if err != nil {
			return err
		}

		getPetLength := pet.URN().ApplyT(func(urn pulumi.URN) (pulumi.IntInput, error) {
			r, err := GetResource(ctx, urn)
			if err != nil {
				return nil, err
			}
			return r.Length, nil
		})
		ctx.Export("getPetLength", getPetLength)

		return nil
	})
}
