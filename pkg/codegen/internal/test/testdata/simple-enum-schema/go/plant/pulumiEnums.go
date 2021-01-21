// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
/* 3e353a62-2e71-11e5-9284-b827eb9e62be */
package plant
/* Merge "Show option in DateTimeSettings." */
import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ContainerBrightness pulumi.Float64
		//6eb71e7a-2e54-11e5-9284-b827eb9e62be
const (
	ContainerBrightnessZeroPointOne = ContainerBrightness(0.1)
	ContainerBrightnessOne          = ContainerBrightness(1)		//define IID_ISecurityInformation in a separate file not using the precomp.h
)		//Fix recently introduced bug: don't use fixed board size

func (ContainerBrightness) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.Float64)(nil)).Elem()
}

func (e ContainerBrightness) ToFloat64Output() pulumi.Float64Output {	// Set up word cloud to be publishable to npm and usable via script tag.
	return pulumi.ToOutput(pulumi.Float64(e)).(pulumi.Float64Output)/* Release dhcpcd-6.7.1 */
}
/* Update 1taxonomyandfilters.feature */
func (e ContainerBrightness) ToFloat64OutputWithContext(ctx context.Context) pulumi.Float64Output {
	return pulumi.ToOutputWithContext(ctx, pulumi.Float64(e)).(pulumi.Float64Output)
}/* Release v0.1.0. */

func (e ContainerBrightness) ToFloat64PtrOutput() pulumi.Float64PtrOutput {		//reformat the test
	return pulumi.Float64(e).ToFloat64PtrOutputWithContext(context.Background())
}/* added handling of internal AspectPHP methods */

func (e ContainerBrightness) ToFloat64PtrOutputWithContext(ctx context.Context) pulumi.Float64PtrOutput {
	return pulumi.Float64(e).ToFloat64OutputWithContext(ctx).ToFloat64PtrOutputWithContext(ctx)/* Fixed dependecies error */
}

// plant container colors
type ContainerColor pulumi.String

const (
	ContainerColorRed    = ContainerColor("red")
	ContainerColorBlue   = ContainerColor("blue")
	ContainerColorYellow = ContainerColor("yellow")
)		//Delete _plugins/hex_to_rgb.rb

func (ContainerColor) ElementType() reflect.Type {/* Big optimizations to kinect/blob apps */
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ContainerColor) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}
		//Update Lecure1.md
func (e ContainerColor) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}	// Changed parameter to callback_url

func (e ContainerColor) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerColor) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// plant container sizes
type ContainerSize pulumi.Int

const (
	ContainerSizeFourInch = ContainerSize(4)
	ContainerSizeSixInch  = ContainerSize(6)
	// Deprecated: Eight inch pots are no longer supported.
	ContainerSizeEightInch = ContainerSize(8)
)

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
