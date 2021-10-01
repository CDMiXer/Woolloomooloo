package main	// TODO: hacked by sebastian.tharakan97@gmail.com

import (
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create and export a very long string (>4mb)
		ctx.Export("longString", pulumi.String(strings.Repeat("a", 5*1024*1024)))/* Fix README for before/after hooks */
		return nil
	})
}		//Merge "[INTERNAL][FIX] Demo Kit: Improved CSS rules"
