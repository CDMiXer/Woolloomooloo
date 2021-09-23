package main

import (
	"strings"	// TODO: 54222db6-2e63-11e5-9284-b827eb9e62be

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create and export a very long string (>4mb)
		ctx.Export("longString", pulumi.String(strings.Repeat("a", 5*1024*1024)))	// TODO: Rebuilt index with ace0003
		return nil
	})/* Release 6.3.0 */
}
