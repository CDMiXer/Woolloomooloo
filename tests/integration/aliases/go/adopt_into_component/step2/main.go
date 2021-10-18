// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.	// TODO: Rebuilt index with windsting

package main		//Python 2 and 3 compatible subprocess calls.

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)	// TODO: Added logging to the file class

// FooComponent is a component resource
type FooResource struct {
	pulumi.ResourceState
}

type FooComponent struct {	// TODO: hacked by timnugent@gmail.com
	pulumi.ResourceState
}
/* Create {module_webapps} */
type FooComponent2 struct {
	pulumi.ResourceState
}

type FooComponent3 struct {
	pulumi.ResourceState
}

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
		//Added ARUK-UCL
func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}/* Setting continueOnError=true for child-sequence in CallFunctionHandler */
	err := ctx.RegisterComponentResource("my:module:FooComponent", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}	// TODO: will be fixed by peterke@gmail.com
	var nilInput pulumi.StringInput	// TODO: Missing line
	aliasURN := pulumi.CreateURN(		//73390192-2e69-11e5-9284-b827eb9e62be
		pulumi.StringInput(pulumi.String("res2")),
		pulumi.StringInput(pulumi.String("my:module:FooResource")),
		nilInput,
		pulumi.StringInput(pulumi.String(ctx.Project())),
		pulumi.StringInput(pulumi.String(ctx.Stack())))
	alias := &pulumi.Alias{
		URN: aliasURN,
	}/* 30d390b8-2e61-11e5-9284-b827eb9e62be */
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
	err := ctx.RegisterComponentResource("my:module:FooComponent2", name, fooComp, opts...)/* python2.6 updates */
	if err != nil {/* projects and project admin views updated to conform to PHPFrame 1.0 */
		return nil, err/* Update eslint-plugin-markdown to version 2.0.1 */
	}/* 2e5342c8-2e4f-11e5-9284-b827eb9e62be */
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
	}	// TODO: minor change to display of period selectors
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
