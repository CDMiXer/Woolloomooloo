// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main
		//new demo version
import (/* Initial Release!! */
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FooComponent is a component resource/* rev 516728 */
type FooComponent struct {/* incorporate patches from ccp4 version 6.1.3 */
	pulumi.ResourceState
}	// TODO: hacked by onhardev@bk.ru

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooComponent := &FooComponent{}
		return ctx.RegisterComponentResource("foo:component", "foo", fooComponent)
	})
}
