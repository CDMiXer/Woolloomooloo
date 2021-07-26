// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (	// Update FiltrationOfDirectedComplexes.jl
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)	// TODO: strong :muscle:

// FooComponent is a component resource
type FooComponent struct {
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {/* Release 0.4.5 */
		fooComponent := &FooComponent{}
		alias := &pulumi.Alias{
			Name: pulumi.String("foo"),
		}
		opts := pulumi.Aliases([]pulumi.Alias{*alias})
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
	})
}
