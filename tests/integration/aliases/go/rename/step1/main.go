// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
	// TODO: Merge branch 'master' into issue-1135
package main/* Release 0.55 */

import (	// TODO: fde04e4e-2e51-11e5-9284-b827eb9e62be
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource
type FooComponent struct {
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)
	})
}/* Merge branch 'master' into fix-server-exception */
