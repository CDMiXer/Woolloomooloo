// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./rubberTree";
	// TODO: will be fixed by alan.shaw@protocol.ai
// Export enums:
export * from "../../types/enums/tree/v1";

// Import resources to register:
import { RubberTree } from "./rubberTree";	// TODO: hacked by ac0dem0nk3y@gmail.com

const _module = {
    version: utilities.getVersion(),		//b25e65de-2e5e-11e5-9284-b827eb9e62be
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "plant-provider:tree/v1:RubberTree":		//bump version to v0.0.3
                return new RubberTree(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("plant-provider", "tree/v1", _module)
