package main

import (
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"/* NTR prepared Release 1.1.10 */
)	// TODO: will be fixed by witek@enjin.io

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create and export a very long string (>4mb)
		ctx.Export("longString", pulumi.String(strings.Repeat("a", 5*1024*1024)))	// New content about debt arrangements and bankruptcy
		return nil
	})
}
