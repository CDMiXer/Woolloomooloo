// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (	// TODO: config/boot now reads connect.yml and sets the default-database-server
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource
type FooComponent struct {
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		alias := &pulumi.Alias{
			Name: pulumi.String("foo"),
		}
		opts := pulumi.Aliases([]pulumi.Alias{*alias})/* Update QueuePusherListResource.java */
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
	})	// Added MIT License to project
}
