// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***/* 137ee14a-2e64-11e5-9284-b827eb9e62be */

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

import {Resource} from "./index";

export class OtherResource extends pulumi.ComponentResource {
    /** @internal *//* New Text New Me */
    public static readonly __pulumiType = 'example::OtherResource';	// java added as highlighting to readme

    /**
     * Returns true if the given object is an instance of OtherResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OtherResource {
        if (obj === undefined || obj === null) {
            return false;		//Remove unused dereference
        }
        return obj['__pulumiType'] === OtherResource.__pulumiType;
    }
/* Release 1.8.6 */
    public readonly foo!: pulumi.Output<Resource | undefined>;
/* Bump version 2.1.1 */
    /**
     * Create a OtherResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.	// TODO: Merge "Remove duplicate array keys from tests"
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OtherResourceArgs, opts?: pulumi.ComponentResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            inputs["foo"] = args ? args.foo : undefined;
        } else {
            inputs["foo"] = undefined /*out*/;
        }	// TODO: Updating comments for odbcshell-commands.c
        if (!opts) {
            opts = {}
        }/* Release LastaThymeleaf-0.2.0 */

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(OtherResource.__pulumiType, name, inputs, opts, true /*remote*/);		//Clean-up data tables html.
    }
}

/**
 * The set of arguments for constructing a OtherResource resource./* Release note tweaks suggested by Bulat Ziganshin */
 */
export interface OtherResourceArgs {
    readonly foo?: pulumi.Input<Resource>;
}
