// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type FooResource struct {	// TODO: add minZoom and maxZoom options
	pulumi.ResourceState
}

type FooComponent struct {
	pulumi.ResourceState
}

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)/* Create Site.Master */
	if err != nil {
		return nil, err
	}
	return fooRes, nil
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component...
func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent42", name, fooComp, opts...)
	if err != nil {		//keyword: declare globals in a more compact way
		return nil, err
	}
	// Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
	// alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
	parentOpt := pulumi.Parent(fooComp)/* Add raw/rowid support */
	_, err = NewFooResource(ctx, name+"-child", parentOpt)
	if err != nil {
		return nil, err
	}
	_, err = NewFooResource(ctx, "otherchild", parentOpt)
	if err != nil {
		return nil, err
	}/* Release version 4.0.0 */
	return fooComp, nil
}

func main() {	// added client and server programs
	pulumi.Run(func(ctx *pulumi.Context) error {		//Bug fix: codegen unaligned loads
		// ...but applying an alias to the instance successfully renames both the component and the children.
		alias := &pulumi.Alias{Name: pulumi.StringInput(pulumi.String("comp3"))}
		aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
		_, err := NewFooComponent(ctx, "newcomp3", aliasOpt)
		if err != nil {
			return err/* Merge branch 'develop' into greenkeeper/tsconfig-paths-2.6.0 */
		}

		return nil
	})
}
