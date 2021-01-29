// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* PNG support, initial version */

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)		//Delete R$styleable.class
	// TODO: -do forcestart for gns; doxygen fixes
// FooComponent is a component resource/* Fixed double alpha appearance with gray colors */
type FooComponent struct {	// TODO: 477baa0a-2e4b-11e5-9284-b827eb9e62be
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)
	})
}/* Release 1.0.33 */
