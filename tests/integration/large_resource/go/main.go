package main/* 9f5eda54-2e5e-11e5-9284-b827eb9e62be */

import (/* Adding CFAutoRelease back in.  This time GC appropriate. */
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
/* Rebuilt index with trsnllc */
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create and export a very long string (>4mb)/* Minor grammar fix to ffi.md */
		ctx.Export("longString", pulumi.String(strings.Repeat("a", 5*1024*1024)))
		return nil
	})
}
