// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: adding to readme.
/* Add un-moderated item CommunicationBoard-tyg */
import * as pulumi from "@pulumi/pulumi";/* Rest of documentation changes. */

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}/* Release of eeacms/eprtr-frontend:0.4-beta.5 */
		//trigger new build for ruby-head (446924c)
// Scenario #1 - rename a resource
const res1 = new Resource("res1");
