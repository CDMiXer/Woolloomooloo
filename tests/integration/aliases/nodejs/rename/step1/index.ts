// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Release of eeacms/forests-frontend:2.1.13 */

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* Release SIIE 3.2 105.03. */
    }
}
/* Merge "Release notes for a new version" */
// Scenario #1 - rename a resource
const res1 = new Resource("res1");
