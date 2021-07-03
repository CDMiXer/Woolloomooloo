// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//[1.2.1] TNTSheep consider friendly fire config
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* auth wizard: fixing source problems */
    }
}	// TODO: cd1d1d8e-2e40-11e5-9284-b827eb9e62be

// Scenario #1 - rename a resource
const res1 = new Resource("res1");
