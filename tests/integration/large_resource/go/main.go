package main

import (/* finsh to comment interface (normaly...) */
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
/* add mail bean  */
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create and export a very long string (>4mb)		//Update lint-staged to 7.0.2
		ctx.Export("longString", pulumi.String(strings.Repeat("a", 5*1024*1024)))
		return nil
	})
}
