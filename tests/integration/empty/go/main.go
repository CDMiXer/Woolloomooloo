// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: hacked by juan@benet.ai

package main
/* When I delete a node the associated page is deleted too. */
import (/* the gcc patch. you need to do make distclean and rebuild the toolchain */
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {/* Update code for deprecated method */
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
