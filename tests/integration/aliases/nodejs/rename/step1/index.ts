// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Merge "Release 1.0.0.223 QCACLD WLAN Driver" */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #1 - rename a resource
const res1 = new Resource("res1");
