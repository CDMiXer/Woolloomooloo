// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* update srv-delivery user, prompt for root password */
    }/* Prepare Release 1.1.6 */
}

// Scenario #1 - rename a resource
const res1 = new Resource("res1");
