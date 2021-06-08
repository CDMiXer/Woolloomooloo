// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main
	// TODO: will be fixed by alan.shaw@protocol.ai
import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"	// TODO: Upgrade rabbitmq
)

// FooComponent is a component resource
type FooComponent struct {
	pulumi.ResourceState/* Add step to include creating a GitHub Release */
}

func main() {	// added testing instructions to readme
	pulumi.Run(func(ctx *pulumi.Context) error {		//Fix context menu offset
		fooComponent := &FooComponent{}
		alias := &pulumi.Alias{
			Name: pulumi.String("foo"),
		}
		opts := pulumi.Aliases([]pulumi.Alias{*alias})
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
	})
}
