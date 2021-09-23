// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// Merge "Add API list=extensions to get the list of known extensions"
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }/* Release of eeacms/eprtr-frontend:0.4-beta.13 */
}

// Scenario #5 - composing #1 and #3
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchildrenamed", {
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],/* Rename Platform::zendClassExists() to Platform::classExists() */
        });
    }		//Updated readme file via branch readme-edits
}
const comp5 = new ComponentFive("newcomp5", {
    aliases: [{ name: "comp5" }],
});
