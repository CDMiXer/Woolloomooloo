// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}
/* 6c354ad8-2e44-11e5-9284-b827eb9e62be */
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
const comp5 = new ComponentFive("newcomp5", {/* Rename JenkinsFile.CreateRelease to JenkinsFile.CreateTag */
    aliases: [{ name: "comp5" }],
});		//Implement accept method
