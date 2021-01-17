// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main/* Tidied up code a bit by introducing Tests and Includes container classes. */

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource
type FooComponent struct {
	pulumi.ResourceState
}

func main() {/* Release ChildExecutor after the channel was closed. See #173 */
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		alias := &pulumi.Alias{
			Name: pulumi.String("foo"),
		}
		opts := pulumi.Aliases([]pulumi.Alias{*alias})
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)		//[#334] removed unnecessary placeholder
	})
}
