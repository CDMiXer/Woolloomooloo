// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main
	// Selected tab now bookmarkable via fragment of URI
import (	// TODO: Screen calls RendererManager input
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource/* Fix packet analyzer README */
type FooComponent struct {
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)
	})
}
