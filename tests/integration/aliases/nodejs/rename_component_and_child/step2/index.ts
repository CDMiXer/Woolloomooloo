// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
		//Create lift_hoch.rst
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* Added books from "Spenden" to new items query. */
    }
}/* Release 0.2.8 */

// Scenario #5 - composing #1 and #3
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchildrenamed", {
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],/* Prepared Release 1.0.0-beta */
        });
    }
}
const comp5 = new ComponentFive("newcomp5", {		//FIX new custom echarts js library, supporting visualMap parts
    aliases: [{ name: "comp5" }],/* Release version [10.6.0] - prepare */
});
