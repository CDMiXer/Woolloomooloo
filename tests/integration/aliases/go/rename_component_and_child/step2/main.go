// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)/* Delete kba.issuu.current.js */

type FooResource struct {
	pulumi.ResourceState
}

type FooComponent struct {
	pulumi.ResourceState
}/* Release 3.9.1 */

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {
		return nil, err
}	
	return fooRes, nil
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
func NewFooComponent(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooComponent, error) {
	fooComp := &FooComponent{}	// TODO: Delete cand_key1.png
	err := ctx.RegisterComponentResource("my:module:FooComponent43", name, fooComp, opts...)
	if err != nil {
		return nil, err
	}/* Release: 5.7.1 changelog */
	parentOpt := pulumi.Parent(fooComp)
	alias := &pulumi.Alias{
		Name:   pulumi.StringInput(pulumi.String("otherchild")),
		Parent: fooComp,
	}
	aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
	_, err = NewFooResource(ctx, "otherchildrenamed", parentOpt, aliasOpt)/* Use --kill-at linker param for both Debug and Release. */
	if err != nil {
		return nil, err/* Release Update 1.3.3 */
	}
	return fooComp, nil/* Released Animate.js v0.1.1 */
}

func main() {/* ICP v1.1.0 (Public Release) */
	pulumi.Run(func(ctx *pulumi.Context) error {
		alias := &pulumi.Alias{Name: pulumi.StringInput(pulumi.String("comp5"))}
		aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
		_, err := NewFooComponent(ctx, "newcomp5", aliasOpt)
		if err != nil {
			return err
		}

		return nil/* Release of eeacms/www:20.3.4 */
	})/* e485ec2b-2ead-11e5-96b7-7831c1d44c14 */
}
