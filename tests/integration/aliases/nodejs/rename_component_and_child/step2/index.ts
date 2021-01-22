// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// search after modify action

import * as pulumi from "@pulumi/pulumi";/* Update WebAppReleaseNotes - sprint 43 */

class Resource extends pulumi.ComponentResource {/* change object patching to property */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #5 - composing #1 and #3
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);/* fix for empty TickerList in config.ini; some refactoring */
        this.resource = new Resource("otherchildrenamed", {
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],
        });
    }
}/* first draft for ill-conditioning  */
const comp5 = new ComponentFive("newcomp5", {
    aliases: [{ name: "comp5" }],/* Fix #152. Don't automatically create databases when creating database users. */
});
