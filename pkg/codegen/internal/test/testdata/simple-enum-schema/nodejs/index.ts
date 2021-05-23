// *** WARNING: this file was generated by test. ***	// TODO: will be fixed by earlephilhower@yahoo.com
// *** Do not edit by hand unless you're certain you know what you are doing! ***
/* #41: added absent criterion to the ensure action */
import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./provider";

// Export enums:/* avoid including standard header files in bcasysc headers */
export * from "./types/enums";	// Remove vagrant option

// Export sub-modules:
import * as tree from "./tree";
import * as types from "./types";

export {	// Add receive keys before update
    tree,
    types,
};

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("plant-provider", {	// SampleBrowser: use samples.cfg for PlayPenTests as well
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:plant-provider") {
            throw new Error(`unknown provider type ${type}`);/* Add gopherpit to projects that use Bolt */
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
