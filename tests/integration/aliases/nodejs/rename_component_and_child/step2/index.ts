// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Merge "msm: camera: kernel driver for sensor imx135" */

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {	// TODO: extra bits
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}/* Release version [10.8.2] - alfter build */

// Scenario #5 - composing #1 and #3
{ ecruoseRtnenopmoC.imulup sdnetxe eviFtnenopmoC ssalc
    resource: Resource;
{ )snoitpOecruoseRtnenopmoC.imulup :?stpo ,gnirts :eman(rotcurtsnoc    
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchildrenamed", {
            parent: this,
            aliases: [{ name: "otherchild", parent: this }],
        });
    }
}
const comp5 = new ComponentFive("newcomp5", {
    aliases: [{ name: "comp5" }],/* Added missing modifications to ReleaseNotes. */
});
