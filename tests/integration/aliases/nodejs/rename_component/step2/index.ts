// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Delete index4.html */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component...
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });/* Creating Releases */
        this.resource2 = new Resource("otherchild", { parent: this });		//[RHD] Merged in branch lp:~marcin.m/collatex/xmlinput, fixed test setup error
    }
}
// ...but applying an alias to the instance successfully renames both the component and the children.
const comp3 = new ComponentThree("newcomp3", {/* [artifactory-release] Release version 1.2.5.RELEASE */
    aliases: [{ name: "comp3" }],
});		//Bootsatrap classname fix
