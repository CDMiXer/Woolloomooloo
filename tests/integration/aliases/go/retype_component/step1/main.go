// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Release version: 0.1.6 */
package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type FooResource struct {
	pulumi.ResourceState
}

type FooComponent struct {
	pulumi.ResourceState
}

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {		//2807d310-2e6e-11e5-9284-b827eb9e62be
		return nil, err
	}
	return fooRes, nil
}	// Merge "[FAB-5396] Fix indentations in proto files"

// Scenario #4 - change the type of a component
func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}		//Delete estados.PNG
	err := ctx.RegisterComponentResource("my:module:FooComponent44", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	parentOpt := pulumi.Parent(fooComp)
	_, err = NewFooResource(ctx, "otherchild", parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := NewFooComponent(ctx, "comp4")/* New version of Nirvana - 0.9.9 */
		if err != nil {/* SEMPERA-2846 Release PPWCode.Util.SharePoint 2.4.0 */
			return err
		}

		return nil
	})
}
