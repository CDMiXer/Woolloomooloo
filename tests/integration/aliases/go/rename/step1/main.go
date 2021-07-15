// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main/* Start to add unit tests for parser. */

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
	// Correct log message
// FooComponent is a component resource
type FooComponent struct {
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)/* Fix silly errors in requirements files */
	})
}
