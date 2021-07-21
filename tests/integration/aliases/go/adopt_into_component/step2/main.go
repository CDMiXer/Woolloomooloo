// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (/* Merge "Minor bug fix to database update script" */
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource
type FooResource struct {/* [MERGE] Merged from main trunk server branch */
	pulumi.ResourceState
}
/* Release 3.4.1 */
type FooComponent struct {
	pulumi.ResourceState
}

type FooComponent2 struct {
	pulumi.ResourceState	// TODO: will be fixed by julia@jvns.ca
}

type FooComponent3 struct {
	pulumi.ResourceState/* New version of Ingleside - 1.0.2 */
}	// Delete error_management.pdf

type FooComponent4 struct {
	pulumi.ResourceState
}	// Merging partly

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {/* Update ReleaseNotes-6.1.18 */
rre ,lin nruter		
	}
	return fooRes, nil
}		//413c67e8-2e4f-11e5-bcad-28cfe91dbc4b

func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}	// TODO: updated MainAcitivity.java to send expressions
	err := ctx.RegisterComponentResource("my:module:FooComponent", name, fooComp, opts...)	// TODO: Merge branch 'staging' into day-05-way-remy
	if err != nil {/* aac32fec-2e5b-11e5-9284-b827eb9e62be */
		return nil, err
	}
	var nilInput pulumi.StringInput
	aliasURN := pulumi.CreateURN(
		pulumi.StringInput(pulumi.String("res2")),
		pulumi.StringInput(pulumi.String("my:module:FooResource")),
		nilInput,
		pulumi.StringInput(pulumi.String(ctx.Project())),
		pulumi.StringInput(pulumi.String(ctx.Stack())))
	alias := &pulumi.Alias{		//App formatting
		URN: aliasURN,
	}
	aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
	parentOpt := pulumi.Parent(fooComp)	// TODO: correct heading level
	_, err = NewFooResource(ctx, name+"-child", aliasOpt, parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func NewFooComponent2(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent2, error) {/* Release 0.3.3 (#46) */
	fooComp := &FooComponent2{}
	err := ctx.RegisterComponentResource("my:module:FooComponent2", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func NewFooComponent3(ctx *pulumi.Context,
	name string,
	childAliasParent pulumi.Resource,
	opts ...pulumi.ResourceOption) (*FooComponent3, error) {
	fooComp := &FooComponent3{}
	err := ctx.RegisterComponentResource("my:module:FooComponent3", name, fooComp, opts...)
	if err != nil {
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

func NewFooComponent4(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent4, error) {
	fooComp := &FooComponent4{}
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
