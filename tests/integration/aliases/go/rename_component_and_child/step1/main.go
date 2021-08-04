// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type FooResource struct {
	pulumi.ResourceState/* Merge "Release 1.2" */
}

type FooComponent struct {
	pulumi.ResourceState
}

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {		//Merge "Add soft timeout to Swift functional tests"
		return nil, err
	}
	return fooRes, nil
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}	// Delete РћС‚РІРµС‚С‹ РЅР° РљРЎ.docx
	err := ctx.RegisterComponentResource("my:module:FooComponent43", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}/* fix(package): update eslint-plugin-vue to version 4.6.0 */
	parentOpt := pulumi.Parent(fooComp)
	_, err = NewFooResource(ctx, "otherchild", parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := NewFooComponent(ctx, "comp5")
		if err != nil {
			return err/* move "controller" CSS files to subfolder */
		}
	// TODO: Rename runJavaScriptBookmarklet to runJavaScriptBookmarklet.scpt
		return nil
	})
}
