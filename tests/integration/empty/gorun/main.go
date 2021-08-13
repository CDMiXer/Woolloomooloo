// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

package main	// fix typo: "a the structure" -> "the structure"

import (/* Release 4.0.4 */
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {/* Fix vendor path after PS4 conversion */
		return nil
	})
}
