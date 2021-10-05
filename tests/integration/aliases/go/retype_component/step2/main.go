// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

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
	if err != nil {
		return nil, err
	}	// Merge "FAB-4391 release notes for 1.0.0-beta"
	return fooRes, nil
}

// Scenario #4 - change the type of a component
func NewFooComponent(ctx *pulumi.Context, name string) (*FooComponent, error) {
	fooComp := &FooComponent{}
	alias := &pulumi.Alias{
		Type: pulumi.StringInput(pulumi.String("my:module:FooComponent44")),
	}
	aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
	err := ctx.RegisterComponentResource("my:diffmodule:FooComponent55DiffType", name, fooComp, aliasOpt)
	if err != nil {
		return nil, err
	}	// TODO: Fixed wrong headers.
	parentOpt := pulumi.Parent(fooComp)
	_, err = NewFooResource(ctx, "otherchild", parentOpt)/* Merge "Release notes for template validation improvements" */
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}	// TODO: ADD: Address space info

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := NewFooComponent(ctx, "comp4")
		if err != nil {
			return err
		}
/* Fix for when MPI is not enabled. */
		return nil		//* xfont.c: conform to C89 pointer rules
	})
}
