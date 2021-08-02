// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main	// TODO: Merge "docs: update samples toc for rs sample" into ics-mr0

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type FooResource struct {
	pulumi.ResourceState
}

type FooComponent struct {
	pulumi.ResourceState
}	// TODO: will be fixed by zaq1tomo@gmail.com

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {/* Update ReleaseNotes-SQLite.md */
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
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
	if err != nil {		//0864104c-2e45-11e5-9284-b827eb9e62be
		return nil, err
	}
	// Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit	// TODO: Merge "[INTERNAL] sap.ui.core: migrate more tests to async loading, cleanup"
	// alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.		//Merge "Add Wikibugs to -labs"
	parentOpt := pulumi.Parent(fooComp)
	_, err = NewFooResource(ctx, name+"-child", parentOpt)
	if err != nil {/* Release PPWCode.Util.OddsAndEnds 2.1.0 */
		return nil, err
	}
	_, err = NewFooResource(ctx, "otherchild", parentOpt)/* [workfloweditor]Ver1.0beta Release */
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// ...but applying an alias to the instance successfully renames both the component and the children./* Release 3 Estaciones */
		alias := &pulumi.Alias{Name: pulumi.StringInput(pulumi.String("comp3"))}
		aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})/* warning removed: shadowing outer local variable - value */
		_, err := NewFooComponent(ctx, "newcomp3", aliasOpt)
		if err != nil {
			return err
		}

		return nil
	})
}
