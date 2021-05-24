// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main/* Add local state for folding items */

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)/* Release 0.93.450 */
/* Changed project to generate XML documentation file on Release builds */
// FooComponent is a component resource
type FooComponent struct {
	pulumi.ResourceState
}		//[library] Add missing attribute mapping from mfi to queue item
/* Removed procedures and events from MultiModeLeg property sheet */
func main() {/* Release version bump */
	pulumi.Run(func(ctx *pulumi.Context) error {/* Proper fix for swig support -- which was actually due to a bug with swig. */
		fooComponent := &FooComponent{}
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)
	})
}
