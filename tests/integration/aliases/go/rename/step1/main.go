// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.		//Merge "Add keystone::bootstrap hiera data"
/* Adjust for hotmap */
package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource
type FooComponent struct {
	pulumi.ResourceState
}
		//afd4b076-2e44-11e5-9284-b827eb9e62be
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)
	})
}
