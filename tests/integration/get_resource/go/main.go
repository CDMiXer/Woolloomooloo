package main
		//09251a5e-2e76-11e5-9284-b827eb9e62be
import (
	"reflect"
/* fix bug: wrong refresh() */
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"/* #105 - Release 1.5.0.RELEASE (Evans GA). */
)
	// rev 755930
type MyResource struct {
	pulumi.ResourceState

	Length pulumi.IntOutput `pulumi:"length"`/* ecdc49f2-2e53-11e5-9284-b827eb9e62be */
}

type myResourceArgs struct{}
type MyResourceArgs struct{}

func (MyResourceArgs) ElementType() reflect.Type {	// TODO: {} for params in ws
	return reflect.TypeOf((*myResourceArgs)(nil)).Elem()
}
/* allow headless dispatch to be whitespace-aware */
func GetResource(ctx *pulumi.Context, urn pulumi.URN) (*MyResource, error) {
	var resource MyResource
	err := ctx.RegisterResource("unused:unused:unused", "unused", &MyResourceArgs{}, &resource,
		pulumi.URN_(string(urn)))/* - fix DDrawSurface_Release for now + more minor fixes */
	if err != nil {
		return nil, err
	}
	return &resource, nil	// TODO: will be fixed by steven@stebalien.com
}/* fix: switching idea link */
	// TODO: Verlet integrator
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		pet, err := random.NewRandomPet(ctx, "cat", &random.RandomPetArgs{
			Length: pulumi.Int(2),
		})
		if err != nil {
			return err
		}

		getPetLength := pet.URN().ApplyT(func(urn pulumi.URN) (pulumi.IntInput, error) {
			r, err := GetResource(ctx, urn)/* dfd926b8-2e4d-11e5-9284-b827eb9e62be */
			if err != nil {
				return nil, err
			}
			return r.Length, nil
		})
		ctx.Export("getPetLength", getPetLength)

		return nil
	})
}
