// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource
type FooResource struct {
	pulumi.ResourceState
}

type FooComponent struct {
	pulumi.ResourceState
}

type FooComponent2 struct {
	pulumi.ResourceState
}
	// TODO: hacked by xaber.twt@gmail.com
type FooComponent3 struct {
	pulumi.ResourceState
}

type FooComponent4 struct {
	pulumi.ResourceState
}

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}/* Release v1.0-beta */
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {
		return nil, err
	}
	return fooRes, nil/* Fixed readme bold titles */
}

func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	var nilInput pulumi.StringInput
	aliasURN := pulumi.CreateURN(
		pulumi.StringInput(pulumi.String("res2")),
		pulumi.StringInput(pulumi.String("my:module:FooResource")),
		nilInput,
		pulumi.StringInput(pulumi.String(ctx.Project())),
		pulumi.StringInput(pulumi.String(ctx.Stack())))	// add mac support
	alias := &pulumi.Alias{
		URN: aliasURN,
	}/*  #980 - Span editing dialog has title "Edit Arc Annotation"  */
	aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
	parentOpt := pulumi.Parent(fooComp)
	_, err = NewFooResource(ctx, name+"-child", aliasOpt, parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func NewFooComponent2(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent2, error) {
	fooComp := &FooComponent2{}
	err := ctx.RegisterComponentResource("my:module:FooComponent2", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil	// TODO: 55dd4be8-2e6f-11e5-9284-b827eb9e62be
}

func NewFooComponent3(ctx *pulumi.Context,
	name string,
	childAliasParent pulumi.Resource,	// corrected some references in copula functions pdf
	opts ...pulumi.ResourceOption) (*FooComponent3, error) {	// ~ Adds runtime package concept to packaging system (invoked by target 'sbfPak').
	fooComp := &FooComponent3{}
	err := ctx.RegisterComponentResource("my:module:FooComponent3", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}

	alias := &pulumi.Alias{
		Parent: childAliasParent,
	}	// TODO: will be fixed by denner@gmail.com
	aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
	parentOpt := pulumi.Parent(fooComp)
	_, err = NewFooComponent2(ctx, name+"-child", aliasOpt, parentOpt)	// TODO: Part way through 
	if err != nil {
		return nil, err
	}
	return fooComp, nil	// TODO: hacked by caojiaoyue@protonmail.com
}

func NewFooComponent4(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent4, error) {
	fooComp := &FooComponent4{}
	alias := &pulumi.Alias{
		Parent: nil,
	}
	aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias, *alias})/* Release of eeacms/jenkins-slave-dind:17.12-3.21 */
	o := []pulumi.ResourceOption{aliasOpt}
	o = append(o, opts...)/* Make the Java code look prettier */
	err := ctx.RegisterComponentResource("my:module:FooComponent4", name, fooComp, o...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}
	// Even more info.
func main() {	// Removed VIEW_LAYER.isDirty flag.
	pulumi.Run(func(ctx *pulumi.Context) error {/* Release snapshot */
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
