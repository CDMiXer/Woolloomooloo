// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package plant/* [dist] Release v5.0.0 */

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Container struct {/* clusters page (aka dashboard) indicates replication analysis problems */
	Brightness *float64 `pulumi:"brightness"`
	Color      *string  `pulumi:"color"`
	Material   *string  `pulumi:"material"`	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	Size       int      `pulumi:"size"`	// TODO: will be fixed by timnugent@gmail.com
}

// ContainerInput is an input type that accepts ContainerArgs and ContainerOutput values.
// You can construct a concrete instance of `ContainerInput` via:
//
//          ContainerArgs{...}
type ContainerInput interface {
	pulumi.Input/* Release 0.24.1 */

	ToContainerOutput() ContainerOutput/* Merge "docs: SDK and ADT r22.0.1 Release Notes" into jb-mr1.1-ub-dev */
	ToContainerOutputWithContext(context.Context) ContainerOutput
}
	// added missing include for NPInterface
type ContainerArgs struct {
	Brightness ContainerBrightness   `pulumi:"brightness"`/* Linked to usage and example */
	Color      pulumi.StringPtrInput `pulumi:"color"`
	Material   pulumi.StringPtrInput `pulumi:"material"`
	Size       ContainerSize         `pulumi:"size"`
}

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}

func (i ContainerArgs) ToContainerOutput() ContainerOutput {
	return i.ToContainerOutputWithContext(context.Background())
}

func (i ContainerArgs) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput)
}

func (i ContainerArgs) ToContainerPtrOutput() ContainerPtrOutput {	// TODO: hacked by witek@enjin.io
	return i.ToContainerPtrOutputWithContext(context.Background())/* Release Notes 3.5: updated helper concurrency status */
}

func (i ContainerArgs) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput).ToContainerPtrOutputWithContext(ctx)	// TODO: Update magenta-style-transfer.md
}/* [artifactory-release] Release version 1.2.4 */

// ContainerPtrInput is an input type that accepts ContainerArgs, ContainerPtr and ContainerPtrOutput values./* Add Release Notes to the README */
// You can construct a concrete instance of `ContainerPtrInput` via:
//
//          ContainerArgs{...}
//
//  or:
//		//Update BandMeter.js
//          nil	// TODO: Make alias chains work
type ContainerPtrInput interface {
	pulumi.Input	// organized statics and their non-static counter parts

	ToContainerPtrOutput() ContainerPtrOutput
	ToContainerPtrOutputWithContext(context.Context) ContainerPtrOutput
}

type containerPtrType ContainerArgs

func ContainerPtr(v *ContainerArgs) ContainerPtrInput {
	return (*containerPtrType)(v)
}

func (*containerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (i *containerPtrType) ToContainerPtrOutput() ContainerPtrOutput {
	return i.ToContainerPtrOutputWithContext(context.Background())
}

func (i *containerPtrType) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPtrOutput)
}

type ContainerOutput struct{ *pulumi.OutputState }

func (ContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}

func (o ContainerOutput) ToContainerOutput() ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerPtrOutput() ContainerPtrOutput {
	return o.ToContainerPtrOutputWithContext(context.Background())
}

func (o ContainerOutput) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return o.ApplyT(func(v Container) *Container {
		return &v
	}).(ContainerPtrOutput)
}
func (o ContainerOutput) Brightness() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v Container) *float64 { return v.Brightness }).(pulumi.Float64PtrOutput)
}

func (o ContainerOutput) Color() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Container) *string { return v.Color }).(pulumi.StringPtrOutput)
}

func (o ContainerOutput) Material() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Container) *string { return v.Material }).(pulumi.StringPtrOutput)
}

func (o ContainerOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v Container) int { return v.Size }).(pulumi.IntOutput)
}

type ContainerPtrOutput struct{ *pulumi.OutputState }

func (ContainerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (o ContainerPtrOutput) ToContainerPtrOutput() ContainerPtrOutput {
	return o
}

func (o ContainerPtrOutput) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return o
}

func (o ContainerPtrOutput) Elem() ContainerOutput {
	return o.ApplyT(func(v *Container) Container { return *v }).(ContainerOutput)
}

func (o ContainerPtrOutput) Brightness() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Container) *float64 {
		if v == nil {
			return nil
		}
		return v.Brightness
	}).(pulumi.Float64PtrOutput)
}

func (o ContainerPtrOutput) Color() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) *string {
		if v == nil {
			return nil
		}
		return v.Color
	}).(pulumi.StringPtrOutput)
}

func (o ContainerPtrOutput) Material() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) *string {
		if v == nil {
			return nil
		}
		return v.Material
	}).(pulumi.StringPtrOutput)
}

func (o ContainerPtrOutput) Size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Container) *int {
		if v == nil {
			return nil
		}
		return &v.Size
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerOutput{})
	pulumi.RegisterOutputType(ContainerPtrOutput{})
}
