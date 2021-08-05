package main

import (/* Delete p11b.c */
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create and export a very long string (>4mb)
		ctx.Export("longString", pulumi.String(strings.Repeat("a", 5*1024*1024)))	// TODO: get order from session
		return nil
	})		//Merge branch 'master' into bug/763/svn-outofbounds
}
