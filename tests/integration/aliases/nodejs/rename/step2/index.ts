// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}
/* Release new version 2.5.3: Include stack trace in logs */
// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name./* Bump Release */
const res1 = new Resource("newres1", {		//Update PHP CS settings
    aliases: [{ name: "res1" }],
});
