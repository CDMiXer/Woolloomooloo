package main

import (		//update description to match current tests
	"strings"/* Even more null checking */

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"	// [#4084873] Reverted to separate Actor and Subject classes
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {/* Fix broken preferences form. */
		// Create and export a very long string (>4mb)
		ctx.Export("longString", pulumi.String(strings.Repeat("a", 5*1024*1024)))
		return nil
	})
}
