// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
		//chore(package): update babel-preset-react-native to version 3.0.2
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//commit before rollback.
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
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit	// TODO: hacked by indexxuan@gmail.com
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.		//also trim the filter-fields
        this.resource1 = new Resource(`${name}-child`, { parent: this });
        this.resource2 = new Resource("otherchild", { parent: this });
    }
}
// ...but applying an alias to the instance successfully renames both the component and the children.
const comp3 = new ComponentThree("newcomp3", {
    aliases: [{ name: "comp3" }],
});
