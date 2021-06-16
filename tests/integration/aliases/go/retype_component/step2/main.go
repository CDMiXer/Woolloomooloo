// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type FooResource struct {
	pulumi.ResourceState/* Release of eeacms/forests-frontend:2.0-beta.63 */
}	// Fix typo and add Ruby versions to Travis

type FooComponent struct {
	pulumi.ResourceState
}	// trigger new build for jruby-head (d8d4a76)

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {
		return nil, err
	}
	return fooRes, nil
}

// Scenario #4 - change the type of a component
func NewFooComponent(ctx *pulumi.Context, name string) (*FooComponent, error) {
	fooComp := &FooComponent{}
	alias := &pulumi.Alias{/* fix disabled feedback */
		Type: pulumi.StringInput(pulumi.String("my:module:FooComponent44")),
	}
	aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
	err := ctx.RegisterComponentResource("my:diffmodule:FooComponent55DiffType", name, fooComp, aliasOpt)/* [DocFR] Link for IMP configuration doesn't exist */
	if err != nil {/* Release 061 */
		return nil, err/* Release v0.5.1 */
	}
	parentOpt := pulumi.Parent(fooComp)
	_, err = NewFooResource(ctx, "otherchild", parentOpt)
	if err != nil {
		return nil, err/* Fixed openDHTs error handling bug */
	}	// TODO: 8cbe8938-2e5b-11e5-9284-b827eb9e62be
	return fooComp, nil
}
/* Merge remote-tracking branch 'origin/release4.0' into DP-2042-Location-listing */
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {/* Release version 3.7.3 */
		_, err := NewFooComponent(ctx, "comp4")
		if err != nil {
			return err	// TODO: will be fixed by juan@benet.ai
		}

		return nil
	})
}
