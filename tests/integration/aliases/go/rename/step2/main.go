// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.	// accessory 4 added: sensor without a track
	// Purge permissions before creating
package main
/* Updated Release Notes. */
import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource
type FooComponent struct {
	pulumi.ResourceState
}
/* Rename PULL_REQUEST_TEMPLATE.md to hello.md */
func main() {	// TODO: will be fixed by why@ipfs.io
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}/* Release updates. */
		alias := &pulumi.Alias{		//22772f60-2e62-11e5-9284-b827eb9e62be
			Name: pulumi.String("foo"),
		}
		opts := pulumi.Aliases([]pulumi.Alias{*alias})
		return ctx.RegisterComponentResource("foo:component", "newfoo", fooComponent, opts)
	})/* d9b69ece-2e4c-11e5-9284-b827eb9e62be */
}
