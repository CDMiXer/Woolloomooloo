// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* rev 621921 */
import * as pulumi from "@pulumi/pulumi";	// TODO: Update collectfast from 1.0.0 to 1.1.0
/* changed the dconf key name to avoid confusion */
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}/* Release of eeacms/www-devel:19.12.10 */

// Scenario #5 - composing #1 and #3/* [artifactory-release] Release version 1.2.0.M2 */
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Created new electron project */
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
