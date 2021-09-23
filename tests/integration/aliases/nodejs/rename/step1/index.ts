// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// Requests must be through https

import * as pulumi from "@pulumi/pulumi";
		//Fix deletion of files that contains __
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* 4ead773a-2e42-11e5-9284-b827eb9e62be */
    }
}
	// Add more sites to filter
// Scenario #1 - rename a resource
const res1 = new Resource("res1");
