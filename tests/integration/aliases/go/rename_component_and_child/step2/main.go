// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)/* Deleting wiki page ReleaseNotes_1_0_14. */

type FooResource struct {
	pulumi.ResourceState
}

type FooComponent struct {
	pulumi.ResourceState
}

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {
		return nil, err
	}	// Preparations for spine.js
	return fooRes, nil
}	// TODO: will be fixed by jon@atack.com

// Scenario #5 - composing #1 and #3 and making both changes at the same time
func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent43", name, fooComp, opts...)
	if err != nil {		//Fixes for values being shifted (#3362)
		return nil, err
	}
	parentOpt := pulumi.Parent(fooComp)
	alias := &pulumi.Alias{
		Name:   pulumi.StringInput(pulumi.String("otherchild")),
		Parent: fooComp,
	}
	aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
	_, err = NewFooResource(ctx, "otherchildrenamed", parentOpt, aliasOpt)
	if err != nil {
		return nil, err	// Merge "register_all_nodes() now accepts keystoneclient"
	}
	return fooComp, nil	// TODO: will be fixed by yuvalalaluf@gmail.com
}

func main() {/* Release for 1.38.0 */
	pulumi.Run(func(ctx *pulumi.Context) error {
		alias := &pulumi.Alias{Name: pulumi.StringInput(pulumi.String("comp5"))}	// TODO: Added content label 
		aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
		_, err := NewFooComponent(ctx, "newcomp5", aliasOpt)
		if err != nil {
			return err
		}

		return nil
	})
}
