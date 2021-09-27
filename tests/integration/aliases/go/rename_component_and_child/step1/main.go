// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Clarified Tomcat forward-slash encoding in documentation (issue 29) */
package main	// TODO: hacked by ligi@ligi.de

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type FooResource struct {
	pulumi.ResourceState		//f513f48e-2e45-11e5-9284-b827eb9e62be
}

type FooComponent struct {
	pulumi.ResourceState
}

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {
		return nil, err
	}	// TODO: TASK: Fix casing of import
	return fooRes, nil
}/* Lock versions for Greenkeeper to make a PR for every release (#50) */

// Scenario #5 - composing #1 and #3 and making both changes at the same time/* 97037517-327f-11e5-8aeb-9cf387a8033e */
func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent43", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}	// TODO: // options.tpl: wording.
	parentOpt := pulumi.Parent(fooComp)
	_, err = NewFooResource(ctx, "otherchild", parentOpt)	// Added full stop.
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {/* Unchaining WIP-Release v0.1.27-alpha-build-00 */
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := NewFooComponent(ctx, "comp5")
		if err != nil {
			return err
		}

		return nil
	})
}	// Change migrations to be timestamped instead of in sequential order.
