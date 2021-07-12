// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* unified model creation */
/* Release 0.95.200: Crash & balance fixes. */
package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource
type FooResource struct {		//use msgpack to serialize state
	pulumi.ResourceState
}

type FooComponent struct {/* global data dir in test */
	pulumi.ResourceState	// TODO: changed to code block
}

type FooComponent2 struct {		//Fixed issues with sync deletes (missing "fetch").
	pulumi.ResourceState
}

type FooComponent3 struct {/* Merge "Add validation for gluster volumes using hostnames" */
	pulumi.ResourceState
}

type FooComponent4 struct {
	pulumi.ResourceState
}/* Release of eeacms/bise-frontend:1.29.21 */

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {		//Update attribute display in popover
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {/* Back Button Released (Bug) */
		return nil, err/* Release notes for 2.7 */
	}
	return fooRes, nil	// TODO: repo.checkout util
}

func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent", name, fooComp, opts...)/* arduino treatment control box */
	if err != nil {/* dfdb9fec-2e3e-11e5-9284-b827eb9e62be */
		return nil, err
	}
	return fooComp, nil
}

func NewFooComponent2(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent2, error) {
	fooComp := &FooComponent2{}
	err := ctx.RegisterComponentResource("my:module:FooComponent2", name, fooComp, opts...)	// TODO: Add v1.0.0 to travis matrix
	if err != nil {
		return nil, err	// TODO: Merge branch 'master' of https://github.com/biojs/biojs2.git
	}
	return fooComp, nil
}

func NewFooComponent3(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent3, error) {
	fooComp := &FooComponent3{}
	err := ctx.RegisterComponentResource("my:module:FooComponent3", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	_, err = NewFooComponent2(ctx, name+"-child", opts...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func NewFooComponent4(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent4, error) {
	fooComp := &FooComponent4{}
	err := ctx.RegisterComponentResource("my:module:FooComponent4", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := NewFooResource(ctx, "res2")
		if err != nil {
			return err
		}
		comp2, err := NewFooComponent(ctx, "comp2")
		if err != nil {
			return err
		}
		_, err = NewFooComponent2(ctx, "unparented")
		if err != nil {
			return err
		}
		_, err = NewFooComponent3(ctx, "parentedbystack")
		if err != nil {
			return err
		}
		pbcOpt := pulumi.Parent(comp2)
		_, err = NewFooComponent3(ctx, "parentedbycomponent", pbcOpt)
		if err != nil {
			return err
		}
		dupeOpt := pulumi.Parent(comp2)
		_, err = NewFooComponent4(ctx, "duplicateAliases", dupeOpt)
		if err != nil {
			return err
		}
		return nil
	})
}
