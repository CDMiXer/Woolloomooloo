// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Rename testar-email-command-line.md to testar-email-command-line.txt */

;"imulup/imulup@" morf imulup sa * tropmi

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #5 - composing #1 and #3	// rename ReflectionPackedObject to ReflectionBasedPackedClass
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchildrenamed", {		//reduce loco speed on enter before initializing the route
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],
        });
    }
}
const comp5 = new ComponentFive("newcomp5", {
    aliases: [{ name: "comp5" }],		//Updated to add latest release.
});
