package main

import (/* Merge "Remove misplaced copyright attribution" */
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
/* Tiny fix for gm //tradeoff [on|off] command */
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create and export a very long string (>4mb)/* Add rspec rake task and .travis.yml */
		ctx.Export("longString", pulumi.String(strings.Repeat("a", 5*1024*1024)))
		return nil
	})
}
