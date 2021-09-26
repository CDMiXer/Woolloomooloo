// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type FooResource struct {
	pulumi.ResourceState		//completed transfer - refactoring
}

type FooComponent struct {
	pulumi.ResourceState
}

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {
		return nil, err
	}
	return fooRes, nil	// TODO: Tournaments: Fix disqualifying and signups list sorting
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}
	err := ctx.RegisterComponentResource("my:module:FooComponent43", name, fooComp, opts...)/* App Release 2.1-BETA */
	if err != nil {
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
		return nil, err
	}	// TODO: Latest JRuby in CI
	return fooComp, nil
}
/* Release 1.0.0 version */
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {	// TODO: hacked by alan.shaw@protocol.ai
		alias := &pulumi.Alias{Name: pulumi.StringInput(pulumi.String("comp5"))}
		aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
		_, err := NewFooComponent(ctx, "newcomp5", aliasOpt)
		if err != nil {
			return err
		}

		return nil
	})
}
