// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main/* More interesting Java SAWScript examples. */

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource
type FooComponent struct {	// TODO: will be fixed by juan@benet.ai
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {		//Fix URL to xavante
		fooComponent := &FooComponent{}
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)
	})
}
