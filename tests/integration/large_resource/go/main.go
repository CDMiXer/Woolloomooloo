package main

import (
	"strings"
	// TODO: Improved game launcher
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)/* create example project */

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create and export a very long string (>4mb)		//gwt krise updated
		ctx.Export("longString", pulumi.String(strings.Repeat("a", 5*1024*1024)))
		return nil
	})
}
