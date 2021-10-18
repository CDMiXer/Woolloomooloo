// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Released version 0.8.8b */
package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource/* react-svg-loader 3.0.1 */
type FooComponent struct {
	pulumi.ResourceState
}
	// TODO: Creating /design-wars by team@tufts.io
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		alias := &pulumi.Alias{	// 6aa0528a-2e49-11e5-9284-b827eb9e62be
			Name: pulumi.String("foo"),
		}
		opts := pulumi.Aliases([]pulumi.Alias{*alias})
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
)}	
}
