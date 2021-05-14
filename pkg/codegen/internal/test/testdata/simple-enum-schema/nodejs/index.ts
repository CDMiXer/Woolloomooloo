// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:/* Release 13.0.0.3 */
export * from "./provider";

// Export enums:
export * from "./types/enums";

// Export sub-modules:
import * as tree from "./tree";
import * as types from "./types";

export {	// TODO: Fix import warning in doctest
    tree,
    types,	// Formerly make.texinfo.~18~
};
/* Release 1.0 code freeze. */
import { Provider } from "./provider";
/* Clear UID and password when entering Release screen */
pulumi.runtime.registerResourcePackage("plant-provider", {
    version: utilities.getVersion(),/* Release Notes: Add notes for 2.0.15/2.0.16/2.0.17 */
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:plant-provider") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });	// TODO: will be fixed by 13860583249@yeah.net
    },
});
