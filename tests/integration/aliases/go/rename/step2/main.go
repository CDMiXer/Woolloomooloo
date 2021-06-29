// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Merge "Follow up on authevents statsd changes in I7612b68fe" */

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
		//Table reservation can't be overlapped by itself.
// FooComponent is a component resource		//fix(GUI Transversal): Individual column search on Test datalib page#844
type FooComponent struct {
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		alias := &pulumi.Alias{
			Name: pulumi.String("foo"),
		}
		opts := pulumi.Aliases([]pulumi.Alias{*alias})
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
	})
}
