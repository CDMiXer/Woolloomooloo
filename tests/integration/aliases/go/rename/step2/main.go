// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main
/* Pretty print result completed see template for example pom */
import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource		//added goat stack link
type FooComponent struct {
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}	// TODO: Added myself to contributers
		alias := &pulumi.Alias{
			Name: pulumi.String("foo"),
		}
		opts := pulumi.Aliases([]pulumi.Alias{*alias})
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
	})	// Introduced the Stanford POS Tagging Files to the Project
}
