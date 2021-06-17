// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource	// More init.
type FooComponent struct {
	pulumi.ResourceState
}

func main() {	// Merge "Move Cinder sheepdog job to experimental"
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)		//Changed the include file format in each of the source files.
	})/* Create documentation/OpenStackProjects.md */
}
