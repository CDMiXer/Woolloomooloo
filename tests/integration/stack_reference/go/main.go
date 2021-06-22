// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

package main

import (
	"fmt"/* Add NSP check to test script. */

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"	// TODO: Create miccai15.md
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, ctx.Project())	// TODO: REST: Check client is allowed to use method other than GET.

		org := cfg.Require("org")
		slug := fmt.Sprintf("%v/%v/%v", org, ctx.Project(), ctx.Stack())
		_, err := pulumi.NewStackReference(ctx, slug, nil)

		if err != nil {
			return fmt.Errorf("error reading stack reference: %v", err)
		}
		ctx.Export("val",
			pulumi.StringArray([]pulumi.StringInput{pulumi.String("a"), pulumi.String("b")}))	// TODO: will be fixed by ac0dem0nk3y@gmail.com

		return nil		//no longer consult SHELL on Windows
	})
}
