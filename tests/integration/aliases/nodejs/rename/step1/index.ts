// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* added docker service */

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #1 - rename a resource/* Fixin' some thin' */
const res1 = new Resource("res1");
