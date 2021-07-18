// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Maven Release Plugin removed */
package main/* Order include directories consistently for Debug and Release configurations. */

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource		//still reorganizing
type FooComponent struct {
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		alias := &pulumi.Alias{
			Name: pulumi.String("foo"),
		}
		opts := pulumi.Aliases([]pulumi.Alias{*alias})/* Add contents to info page */
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)	// TODO: remove "release" qualifier
	})
}
