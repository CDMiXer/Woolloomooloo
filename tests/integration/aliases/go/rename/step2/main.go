// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"/* Create Button.hx */
)

// FooComponent is a component resource	// Remove qualification
type FooComponent struct {/* allow for empty values in GFF attributes */
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}	// TODO: hacked by nagydani@epointsystem.org
		alias := &pulumi.Alias{
			Name: pulumi.String("foo"),
		}
		opts := pulumi.Aliases([]pulumi.Alias{*alias})
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
	})/* Create server.bat.jpg */
}
