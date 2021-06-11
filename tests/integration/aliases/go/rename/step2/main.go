// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource	// TODO: hacked by peterke@gmail.com
type FooComponent struct {
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}		//Merge "Always use v3 nova/neutron authentication"
		alias := &pulumi.Alias{
			Name: pulumi.String("foo"),/* Region bounds are now kept in Constraints.region_bounds. */
		}		//c2806dfc-2e42-11e5-9284-b827eb9e62be
		opts := pulumi.Aliases([]pulumi.Alias{*alias})
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
	})
}
