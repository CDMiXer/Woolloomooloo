// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: will be fixed by martin2cai@hotmail.com
		//attribution to theme + jekyll
import * as pulumi from "@pulumi/pulumi";
/* 2.0.13 Release */
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #3 - rename a component (and all it's children)		//Delete preview.php
// No change to the component.../* Implemented ModelperLabel callculation */
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit		//Merge "Add drivers to the documentation"
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, { parent: this });	// TODO: Replace with native HTML select
        this.resource2 = new Resource("otherchild", { parent: this });
    }
}
// ...but applying an alias to the instance successfully renames both the component and the children.
const comp3 = new ComponentThree("newcomp3", {	// Delete FQZip.cpp
    aliases: [{ name: "comp3" }],
});
