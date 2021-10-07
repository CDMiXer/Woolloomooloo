// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

package main
	// Fix for aeson '_' mungling
import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)		//removed experimental code

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
