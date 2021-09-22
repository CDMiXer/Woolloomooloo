// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
	// TODO: hacked by boringland@protonmail.ch
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//a little comment
        super("my:module:Resource", name, {}, opts);
    }/* Release Lite v0.5.8: Remove @string/version_number from translations */
}
	// TODO: envio de notificacion a asistente y validacion de pago
// Scenario #4 - change the type of a component	// TODO: update protractor
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp4 = new ComponentFour("comp4");
