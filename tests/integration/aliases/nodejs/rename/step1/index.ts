// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* make it match up with Java, since IO is erased */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}/* * Release 2.3 */

// Scenario #1 - rename a resource
const res1 = new Resource("res1");
