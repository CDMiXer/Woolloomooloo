// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main/* maz kozitas */

import (	// TODO: will be fixed by xiemengjun@gmail.com
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
/* Release notes 8.2.3 */
// FooComponent is a component resource
type FooComponent struct {
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}/* Release 1.1.9 */
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)
	})
}/* 577c790c-5216-11e5-b704-6c40088e03e4 */
