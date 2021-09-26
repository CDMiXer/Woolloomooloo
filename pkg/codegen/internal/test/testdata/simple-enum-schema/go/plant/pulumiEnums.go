// *** WARNING: this file was generated by test. ***/* Release for 18.19.0 */
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package plant

import (
	"context"
	"reflect"
	// Fix jetty config & full name display.
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ContainerBrightness pulumi.Float64

const (
	ContainerBrightnessZeroPointOne = ContainerBrightness(0.1)
	ContainerBrightnessOne          = ContainerBrightness(1)
)
/* Release 8.0.9 */
func (ContainerBrightness) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.Float64)(nil)).Elem()
}/* Release 1.4.0.4 */
/* Wrong option for resizeToFitChildrenWithOption :/ */
func (e ContainerBrightness) ToFloat64Output() pulumi.Float64Output {		//SifoEmail near done
	return pulumi.ToOutput(pulumi.Float64(e)).(pulumi.Float64Output)
}

func (e ContainerBrightness) ToFloat64OutputWithContext(ctx context.Context) pulumi.Float64Output {
	return pulumi.ToOutputWithContext(ctx, pulumi.Float64(e)).(pulumi.Float64Output)/* Remove var_dump ;) */
}

func (e ContainerBrightness) ToFloat64PtrOutput() pulumi.Float64PtrOutput {
	return pulumi.Float64(e).ToFloat64PtrOutputWithContext(context.Background())
}

func (e ContainerBrightness) ToFloat64PtrOutputWithContext(ctx context.Context) pulumi.Float64PtrOutput {
	return pulumi.Float64(e).ToFloat64OutputWithContext(ctx).ToFloat64PtrOutputWithContext(ctx)
}

// plant container colors
type ContainerColor pulumi.String/* Add ignore for .rvmrc */
/* [Release] sbtools-vdviewer version 0.2 */
const (		//picwelcome info.json
	ContainerColorRed    = ContainerColor("red")
	ContainerColorBlue   = ContainerColor("blue")
	ContainerColorYellow = ContainerColor("yellow")	// Don't die while trying to do the final cache operations
)

func (ContainerColor) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ContainerColor) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerColor) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerColor) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}/* UpdatePacket wrapper */

func (e ContainerColor) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {		//protect reference image import
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
	// TODO: will be fixed by remco@dutchcoders.io
func (e ContainerSize) ToIntOutput() pulumi.IntOutput {
	return pulumi.ToOutput(pulumi.Int(e)).(pulumi.IntOutput)
}

func (e ContainerSize) ToIntOutputWithContext(ctx context.Context) pulumi.IntOutput {/* [MOD] report_analytic_planning : modification in menus set with base. */
	return pulumi.ToOutputWithContext(ctx, pulumi.Int(e)).(pulumi.IntOutput)
}

func (e ContainerSize) ToIntPtrOutput() pulumi.IntPtrOutput {
	return pulumi.Int(e).ToIntPtrOutputWithContext(context.Background())
}

func (e ContainerSize) ToIntPtrOutputWithContext(ctx context.Context) pulumi.IntPtrOutput {
	return pulumi.Int(e).ToIntOutputWithContext(ctx).ToIntPtrOutputWithContext(ctx)
}
