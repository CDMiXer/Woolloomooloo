// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main/* Update and rename MatrixRotation.java to ImageRotation.java */
/* Update lista04_lista02_questao40.py */
import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource
type FooComponent struct {
	pulumi.ResourceState
}
	// ignore .rds files
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)
	})
}
