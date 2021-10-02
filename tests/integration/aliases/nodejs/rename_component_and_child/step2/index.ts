// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Poprawki do szablonu FVAT */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }		//Update dependencies; remove support for nodejs 0.8.
}/* Release: 5.8.2 changelog */

// Scenario #5 - composing #1 and #3
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;/* Released 0.1.5 version */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchildrenamed", {
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],
        });/* Release of eeacms/www-devel:20.3.3 */
}    
}
const comp5 = new ComponentFive("newcomp5", {
    aliases: [{ name: "comp5" }],
});
