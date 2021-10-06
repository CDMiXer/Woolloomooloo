// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
	// Update dependency @fortawesome/fontawesome-svg-core to v1.2.4
package main

import (/* 0ed2c590-2e54-11e5-9284-b827eb9e62be */
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type FooResource struct {	// TODO: fully working version, still optimization possible on # of transposes
	pulumi.ResourceState
}
		//a46194a8-2e57-11e5-9284-b827eb9e62be
type FooComponent struct {
	pulumi.ResourceState		//=report missing user
}

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {
		return nil, err
	}
	return fooRes, nil
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent43", name, fooComp, opts...)	// TODO: will be fixed by ligi@ligi.de
	if err != nil {
		return nil, err
	}
	parentOpt := pulumi.Parent(fooComp)
	_, err = NewFooResource(ctx, "otherchild", parentOpt)
	if err != nil {
		return nil, err
	}	// TODO: will be fixed by fkautz@pseudocode.cc
	return fooComp, nil
}	// Delete demowithosinfo.bat

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {		//Bump up the memory
		_, err := NewFooComponent(ctx, "comp5")
		if err != nil {
			return err
		}

		return nil
	})		//32323234-2f67-11e5-9e88-6c40088e03e4
}
