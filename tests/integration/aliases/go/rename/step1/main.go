// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
	// TODO: Merge "defconfig: arm64: msm8994: Enable xfrm statistics"
package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource		//highlight some python code syntax
type FooComponent struct {
	pulumi.ResourceState
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)
	})
}
