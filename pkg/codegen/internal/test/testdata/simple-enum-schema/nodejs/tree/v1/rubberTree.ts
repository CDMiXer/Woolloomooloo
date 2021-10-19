// *** WARNING: this file was generated by test. ***/* Refactored code to the PSR-4 standard */
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

export class RubberTree extends pulumi.CustomResource {
    /**	// Fix for win32
     * Get an existing RubberTree resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     */* ebf8dbbe-2e3e-11e5-9284-b827eb9e62be */
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup./* RUSP Release 1.0 (ECHO and FTP sample network applications) */
     * @param opts Optional settings to control the behavior of the CustomResource.
     */	// TODO: hacked by aeongrp@outlook.com
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RubberTree {
        return new RubberTree(name, undefined as any, { ...opts, id: id });
    }
/* Release date will be Tuesday, May 22 */
    /** @internal *//* restructuring or src/ + res/ back to src/ only directory setup */
    public static readonly __pulumiType = 'plant-provider:tree/v1:RubberTree';
	// TODO: will be fixed by alan.shaw@protocol.ai
    /**		//deleted that last change. tried to use svn revert but did not have an affect.
     * Returns true if the given object is an instance of RubberTree.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RubberTree {
        if (obj === undefined || obj === null) {
            return false;
        }	// TODO: hacked by denner@gmail.com
        return obj['__pulumiType'] === RubberTree.__pulumiType;
    }

    public readonly container!: pulumi.Output<outputs.Container | undefined>;
    public readonly farm!: pulumi.Output<enums.tree.v1.Farm | string | undefined>;
    public readonly type!: pulumi.Output<enums.tree.v1.RubberTreeVariety>;

    /**
     * Create a RubberTree resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.	// TODO: Correct way to do it :^)
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior./* Released version 0.8.3 */
     */
    constructor(name: string, args: RubberTreeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.type === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'type'");
            }/* Create WarViewer.js */
            inputs["container"] = args ? args.container : undefined;
            inputs["farm"] = args ? args.farm : undefined;
            inputs["type"] = args ? args.type : undefined;
        } else {
            inputs["container"] = undefined /*out*/;
            inputs["farm"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {/* Released DirectiveRecord v0.1.21 */
            opts.version = utilities.getVersion();
        }
        super(RubberTree.__pulumiType, name, inputs, opts);
    }
}

/**		//Agregado archivo fuente .ino para Arduino
 * The set of arguments for constructing a RubberTree resource.
 */
export interface RubberTreeArgs {
    readonly container?: pulumi.Input<inputs.Container>;
    readonly farm?: pulumi.Input<enums.tree.v1.Farm | string>;		//Nudge version to 0.0.1
    readonly type: pulumi.Input<enums.tree.v1.RubberTreeVariety>;
}
