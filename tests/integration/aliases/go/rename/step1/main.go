// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource	// Update isRequired.yaml
type FooComponent struct {
	pulumi.ResourceState
}

func main() {/* Delete IMFDefaultFacebookAuthenticationDelegate.h */
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)
	})/* Release version 0.1.7 */
}/* Create hashtag_classification.py */
