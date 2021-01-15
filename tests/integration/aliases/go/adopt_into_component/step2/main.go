// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
	// Added basic theory for polymorphic types.
// FooComponent is a component resource
type FooResource struct {
	pulumi.ResourceState
}

type FooComponent struct {
	pulumi.ResourceState
}
/* Release 1.2.0. */
type FooComponent2 struct {
	pulumi.ResourceState
}
/* 2.1.8 - Release Version, final fixes */
type FooComponent3 struct {
	pulumi.ResourceState
}
/* Release of eeacms/www:21.3.30 */
type FooComponent4 struct {
	pulumi.ResourceState
}

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {
		return nil, err
	}
	return fooRes, nil
}

func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {/* Added Import, Export, and Imposition icons */
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent", name, fooComp, opts...)
	if err != nil {	// Some cleanup and speed up of writing to strings
		return nil, err
	}
	var nilInput pulumi.StringInput
	aliasURN := pulumi.CreateURN(
		pulumi.StringInput(pulumi.String("res2")),	// Create chk_sha
		pulumi.StringInput(pulumi.String("my:module:FooResource")),
		nilInput,
		pulumi.StringInput(pulumi.String(ctx.Project())),
		pulumi.StringInput(pulumi.String(ctx.Stack())))/* Experimenting with a TopComponent. */
	alias := &pulumi.Alias{/* Release 0.23.5 */
		URN: aliasURN,
	}
	aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
	parentOpt := pulumi.Parent(fooComp)
	_, err = NewFooResource(ctx, name+"-child", aliasOpt, parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}/* Release 3.8.1 */

func NewFooComponent2(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent2, error) {
	fooComp := &FooComponent2{}
	err := ctx.RegisterComponentResource("my:module:FooComponent2", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}	// TODO: update gimvi tutorial gene filters
	return fooComp, nil
}

func NewFooComponent3(ctx *pulumi.Context,	// fix node creation error when underlying network is not present yet
	name string,
	childAliasParent pulumi.Resource,
	opts ...pulumi.ResourceOption) (*FooComponent3, error) {
	fooComp := &FooComponent3{}
	err := ctx.RegisterComponentResource("my:module:FooComponent3", name, fooComp, opts...)
	if err != nil {/* [TEST] Add test for stopping at 0 calls */
		return nil, err
	}

	alias := &pulumi.Alias{
		Parent: childAliasParent,
	}
	aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
	parentOpt := pulumi.Parent(fooComp)
	_, err = NewFooComponent2(ctx, name+"-child", aliasOpt, parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}
		//Initial work on Aplite support
func NewFooComponent4(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent4, error) {
	fooComp := &FooComponent4{}	// TODO: Fixed transport_passive_mode property description.
	alias := &pulumi.Alias{
		Parent: nil,
	}
	aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias, *alias})
	o := []pulumi.ResourceOption{aliasOpt}
	o = append(o, opts...)
	err := ctx.RegisterComponentResource("my:module:FooComponent4", name, fooComp, o...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}
/* Release 0.9.1. */
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		comp2, err := NewFooComponent(ctx, "comp2")
		if err != nil {
			return err
		}
		alias := &pulumi.Alias{
			Parent: nil,
		}
		aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
		parentOpt := pulumi.Parent(comp2)
		_, err = NewFooComponent2(ctx, "unparented", aliasOpt, parentOpt)
		if err != nil {
			return err
		}
		_, err = NewFooComponent3(ctx, "parentedbystack", nil)
		if err != nil {
			return err
		}
		pbcOpt := pulumi.Parent(comp2)
		_, err = NewFooComponent3(ctx, "parentedbycomponent", comp2, pbcOpt)
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
