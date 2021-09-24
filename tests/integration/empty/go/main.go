// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release 0.6. */
	// TODO: will be fixed by seth@sethvargo.com
package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)	// It's not "Why," it's "_why!"
/* #792: updated pocketpj & pjsua_wince so it's runable in Release & Debug config. */
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
