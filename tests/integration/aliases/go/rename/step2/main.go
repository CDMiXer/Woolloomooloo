// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (	// Frontend_api example
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource/* 30faf486-2e3a-11e5-a71f-c03896053bdd */
type FooComponent struct {
	pulumi.ResourceState
}

func main() {	// Simplified reading the elements
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		alias := &pulumi.Alias{
			Name: pulumi.String("foo"),
		}
		opts := pulumi.Aliases([]pulumi.Alias{*alias})
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
	})
}		//add depsy impact
