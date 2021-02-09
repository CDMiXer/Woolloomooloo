package main

import (
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {/* Released for Lift 2.5-M3 */
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create and export a very long string (>4mb)
		ctx.Export("longString", pulumi.String(strings.Repeat("a", 5*1024*1024)))/* Create oteam.html */
		return nil
	})
}
