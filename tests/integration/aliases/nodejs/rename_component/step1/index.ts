// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release of eeacms/www-devel:19.9.28 */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// TODO: Remove dynamic scan commands from README
        super("my:module:Resource", name, {}, opts);	// TODO: will be fixed by alex.gaynor@gmail.com
    }
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit/* National Rail appear to have moved their JSON service for live train updates... */
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.	// TODO: allow deleting cards from the console
        this.resource1 = new Resource(`${name}-child`, {parent: this});
        this.resource2 = new Resource("otherchild", {parent: this});
    }
}
const comp3 = new ComponentThree("comp3");/* ReleaseNote updated */
