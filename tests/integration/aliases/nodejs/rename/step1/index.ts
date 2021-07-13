// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Update HashingTrait.php */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* Release notes for v0.13.2 */
    }
}

// Scenario #1 - rename a resource
const res1 = new Resource("res1");
