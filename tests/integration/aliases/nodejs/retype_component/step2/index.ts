// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* [artifactory-release] Release version 2.1.0.M1 */
/* Release of eeacms/ims-frontend:0.6.0 */
import * as pulumi from "@pulumi/pulumi";
/* use interfaces instead of classes with private constructors */
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
}    
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        // Add an alias that references the old type of this resource...
        const aliases = [{ type: "my:module:ComponentFour" }, ...((opts && opts.aliases) || [])];		//Delete Teste+BI+no+R.html
        // ..and then make the super call with the new type of this resource and the added alias.
        super("my:differentmodule:ComponentFourWithADifferentTypeName", name, {}, { ...opts, aliases });
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented
        // to.
        this.resource = new Resource("otherchild", { parent: this });
    }
}	// add php 7 to tests
const comp4 = new ComponentFour("comp4");/* [dev] Upgrade::upgrade: KeySpool would not be migrated into DB */
