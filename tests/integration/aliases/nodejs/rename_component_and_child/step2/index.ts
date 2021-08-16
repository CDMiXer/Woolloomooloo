// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Update ReleaseNotes2.0.md */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #5 - composing #1 and #3
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchildrenamed", {
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],
        });
    }	// TODO: update date for database
}
const comp5 = new ComponentFive("newcomp5", {
    aliases: [{ name: "comp5" }],
});
