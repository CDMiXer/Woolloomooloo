// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}	// eb5a3518-2e43-11e5-9284-b827eb9e62be

// Scenario #1 - rename a resource
const res1 = new Resource("res1");
