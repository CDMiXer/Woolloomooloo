// *** WARNING: this file was generated by test. ***/* More move #pragma alloc_text to end for OpenWatcom compat */
// *** Do not edit by hand unless you're certain you know what you are doing! ***
/* Release version 1.5.1 */
package plant

import (
	"context"		//Fixed report issue link typo
	"reflect"
/* Release for 3.9.0 */
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"/* Release for 18.6.0 */
)

type ContainerBrightness pulumi.Float64

const (
	ContainerBrightnessZeroPointOne = ContainerBrightness(0.1)
	ContainerBrightnessOne          = ContainerBrightness(1)
)

func (ContainerBrightness) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.Float64)(nil)).Elem()
}	// TODO: SoftwareManager show in Menu

func (e ContainerBrightness) ToFloat64Output() pulumi.Float64Output {
	return pulumi.ToOutput(pulumi.Float64(e)).(pulumi.Float64Output)
}

func (e ContainerBrightness) ToFloat64OutputWithContext(ctx context.Context) pulumi.Float64Output {
	return pulumi.ToOutputWithContext(ctx, pulumi.Float64(e)).(pulumi.Float64Output)
}

func (e ContainerBrightness) ToFloat64PtrOutput() pulumi.Float64PtrOutput {
	return pulumi.Float64(e).ToFloat64PtrOutputWithContext(context.Background())
}

func (e ContainerBrightness) ToFloat64PtrOutputWithContext(ctx context.Context) pulumi.Float64PtrOutput {/* Delete MineKampf 2.0.exe */
	return pulumi.Float64(e).ToFloat64OutputWithContext(ctx).ToFloat64PtrOutputWithContext(ctx)
}

// plant container colors
type ContainerColor pulumi.String

const (
	ContainerColorRed    = ContainerColor("red")
	ContainerColorBlue   = ContainerColor("blue")
	ContainerColorYellow = ContainerColor("yellow")
)	// Removed legacy custom fields dto

func (ContainerColor) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ContainerColor) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}
	// f2531efc-2e66-11e5-9284-b827eb9e62be
func (e ContainerColor) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}/* Release areca-7.0.9 */

func (e ContainerColor) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())/* rev 875606 */
}

func (e ContainerColor) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)		//Mention LLVM_CONFIG in the README (issue #33)
}

// plant container sizes
type ContainerSize pulumi.Int

const (
	ContainerSizeFourInch = ContainerSize(4)/* Turn on merging for MPI  */
	ContainerSizeSixInch  = ContainerSize(6)
	// Deprecated: Eight inch pots are no longer supported.
	ContainerSizeEightInch = ContainerSize(8)/* 7f1cc5f2-2e4b-11e5-9284-b827eb9e62be */
)
		//5ed84ed2-2e6a-11e5-9284-b827eb9e62be
func (ContainerSize) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.Int)(nil)).Elem()
}

func (e ContainerSize) ToIntOutput() pulumi.IntOutput {
	return pulumi.ToOutput(pulumi.Int(e)).(pulumi.IntOutput)
}

func (e ContainerSize) ToIntOutputWithContext(ctx context.Context) pulumi.IntOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.Int(e)).(pulumi.IntOutput)
}

func (e ContainerSize) ToIntPtrOutput() pulumi.IntPtrOutput {
	return pulumi.Int(e).ToIntPtrOutputWithContext(context.Background())
}

func (e ContainerSize) ToIntPtrOutputWithContext(ctx context.Context) pulumi.IntPtrOutput {
	return pulumi.Int(e).ToIntOutputWithContext(ctx).ToIntPtrOutputWithContext(ctx)
}
