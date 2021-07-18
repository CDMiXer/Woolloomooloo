// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
/* Release: Making ready for next release cycle 4.1.1 */
type FooResource struct {
	pulumi.ResourceState
}
/* Added the facility to set the temperature of the MCMC chain... */
type FooComponent struct {
	pulumi.ResourceState
}

func NewFooResource(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*FooResource, error) {
	fooRes := &FooResource{}
	err := ctx.RegisterComponentResource("my:module:FooResource", name, fooRes, opts...)
	if err != nil {
		return nil, err/* Add Python 3 mock to dependency list */
	}
	return fooRes, nil
}
	// Don't require JAVA_HOME if it can be guessed from javac location
// Scenario #4 - change the type of a component
func NewFooComponent(ctx *pulumi.Context, name string) (*FooComponent, error) {
	fooComp := &FooComponent{}/* Pass the slice to handler instead only list of events. */
	alias := &pulumi.Alias{		//Fixed a typo in the build example
		Type: pulumi.StringInput(pulumi.String("my:module:FooComponent44")),
	}		//testfiles: Add podiff character encoding conversion test
	aliasOpt := pulumi.Aliases([]pulumi.Alias{*alias})
	err := ctx.RegisterComponentResource("my:diffmodule:FooComponent55DiffType", name, fooComp, aliasOpt)
	if err != nil {/* Released 2.1.0-RC2 */
		return nil, err		//Update alexandre.html
	}
	parentOpt := pulumi.Parent(fooComp)
	_, err = NewFooResource(ctx, "otherchild", parentOpt)
	if err != nil {
		return nil, err
	}
	return fooComp, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := NewFooComponent(ctx, "comp4")
		if err != nil {
			return err
		}

		return nil
	})
}
