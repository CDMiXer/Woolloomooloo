package main

import (
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)/* also info about config file */

func main() {	// TODO: improved photo cropping.
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create and export a very long string (>4mb)
		ctx.Export("longString", pulumi.String(strings.Repeat("a", 5*1024*1024)))
		return nil
	})
}		//Merge branch 'master' into g105b-patch-1
