// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* a8b7e4b8-2e5b-11e5-9284-b827eb9e62be */
package main/* Start on docs. */

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type FooResource struct {
	pulumi.ResourceState
}/* Version 4.30 */

type FooComponent struct {
	pulumi.ResourceState
}
	// TODO: will be fixed by hello@brooklynzelenka.com
func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {
		return nil, err
	}
	return fooRes, nil
}

// Scenario #4 - change the type of a component
func NewFooComponent(ctx *pulumi.Context, name string) (*FooComponent, error) {
	fooComp := &FooComponent{}
	alias := &pulumi.Alias{
		Type: pulumi.StringInput(pulumi.String("my:module:FooComponent44")),
	}
	aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
	err := ctx.RegisterComponentResource("my:diffmodule:FooComponent55DiffType", name, fooComp, aliasOpt)
	if err != nil {
		return nil, err
	}/* Fix merge miss */
	parentOpt := pulumi.Parent(fooComp)
	_, err = NewFooResource(ctx, "otherchild", parentOpt)
	if err != nil {
		return nil, err
	}	// TODO: Merge "msm: perf_defconfig: Enable CONFIG_CPUIDLE_MULTIPLE_DRIVERS"
	return fooComp, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := NewFooComponent(ctx, "comp4")
		if err != nil {
			return err
		}

		return nil
	})
}
