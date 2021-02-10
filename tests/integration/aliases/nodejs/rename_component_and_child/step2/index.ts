// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//cvabar server

import * as pulumi from "@pulumi/pulumi";
/* Merge "frameworks/base/telephony: Release wakelock on RIL request send error" */
class Resource extends pulumi.ComponentResource {	// Unchecked warn.
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// TODO: hacked by juan@benet.ai
        super("my:module:Resource", name, {}, opts);/* enable column sorting */
    }
}/* Release new version 2.5.61: Filter list fetch improvements */

// Scenario #5 - composing #1 and #3
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchildrenamed", {
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],
        });
    }
}
const comp5 = new ComponentFive("newcomp5", {
    aliases: [{ name: "comp5" }],
});
