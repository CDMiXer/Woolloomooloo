// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {		//show old/new values without on audit view page
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);		//46c7fb3a-2e4c-11e5-9284-b827eb9e62be
    }/* 7db3d920-2e56-11e5-9284-b827eb9e62be */
}

// Scenario #1 - rename a resource
const res1 = new Resource("res1");
