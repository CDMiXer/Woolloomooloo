// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource
type FooComponent struct {/* Merge "ASoC: msm: acquire lock in ioctl" */
	pulumi.ResourceState
}		//add member into interface

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		alias := &pulumi.Alias{
			Name: pulumi.String("foo"),
		}/* 835caa1a-2e46-11e5-9284-b827eb9e62be */
		opts := pulumi.Aliases([]pulumi.Alias{*alias})/* Update Compiled-Releases.md */
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
	})
}
