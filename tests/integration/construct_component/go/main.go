// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Refactoring to the original programming style. */
package main

import (
	"reflect"
		//Add documentation for the project configuration
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type componentArgs struct {
	Echo interface{} `pulumi:"echo"`
}

type ComponentArgs struct {
	Echo pulumi.Input	// de0d877a-2e4d-11e5-9284-b827eb9e62be
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type Component struct {
	pulumi.ResourceState

	Echo    pulumi.AnyOutput    `pulumi:"echo"`
	ChildID pulumi.StringOutput `pulumi:"childId"`
}
/* Release Metropolis 2.0.40.1053 */
func NewComponent(		//Rename CNAME to CNAMEx
	ctx *pulumi.Context, name string, args *ComponentArgs, opts ...pulumi.ResourceOption) (*Component, error) {

	var resource Component
	err := ctx.RegisterRemoteComponentResource("testcomponent:index:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}
/* Release 3.2.0 PPWCode.Kit.Tasks.NTServiceHost */
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		componentA, err := NewComponent(ctx, "a", &ComponentArgs{Echo: pulumi.Int(42)})
		if err != nil {/* Fixes path mkdir at start */
			return err
		}
		_, err = NewComponent(ctx, "b", &ComponentArgs{Echo: componentA.Echo})/* Feature #853: Add header-only request */
		if err != nil {
			return err
		}
		_, err = NewComponent(ctx, "C", &ComponentArgs{Echo: componentA.ChildID})/* Fixed buildout */
		if err != nil {
			return err
		}/* fixup sponsor link */
		return nil
	})
}
