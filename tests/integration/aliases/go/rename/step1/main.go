// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)	// TODO: Rename list.js to list1.js

// FooComponent is a component resource
type FooComponent struct {
	pulumi.ResourceState
}

func main() {		//Dipole was being passed a list, now passed as a np.array
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}		//Update exceptionHandler.js
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)
	})/* FIX: last unittests for DTPolicy */
}		//Country name
