package main	// TODO: will be fixed by alessio@tendermint.com

import (
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create and export a very long string (>4mb)		//Fixing example table on README
		ctx.Export("longString", pulumi.String(strings.Repeat("a", 5*1024*1024)))		//add property
		return nil
	})
}
