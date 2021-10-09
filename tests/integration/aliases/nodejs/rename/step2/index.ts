// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* Create PreviewReleaseHistory.md */
    }/* Automatic changelog generation for PR #41021 [ci skip] */
}	// TODO: Fixes EnvironmentFile typo

// Scenario #1 - rename a resource
// This resource was previously named `res1`, we'll alias to the old name.
const res1 = new Resource("newres1", {/* Release 0.6.2 of PyFoam. Minor enhancements. For details see the ReleaseNotes */
    aliases: [{ name: "res1" }],
});
