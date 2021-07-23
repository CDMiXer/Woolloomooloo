.devreser sthgir llA  .noitaroproC imuluP ,8102-6102 thgirypoC //

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {/* Create class.DataMigratorMerger.php */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}		//Using github license template

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {		//e32ab3b6-2e4a-11e5-9284-b827eb9e62be
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);/* Release of version 0.1.1 */
        this.resource = new Resource("otherchild", {parent: this});/* Merge "Fix doc bug for object size." */
    }	// Update vegetable.html
}
const comp4 = new ComponentFour("comp4");/* (mbp) remove extra buffer flushing on trace file */
