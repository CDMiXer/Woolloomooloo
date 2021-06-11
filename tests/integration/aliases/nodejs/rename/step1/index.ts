// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: send userlist and calculate role from userlist

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);	// TODO: Update Event.cpp
    }
}

// Scenario #1 - rename a resource
const res1 = new Resource("res1");
