// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

export class RubberTree extends pulumi.CustomResource {
    /**	// Create resp_linear_regression_moesio.ipynb
     * Get an existing RubberTree resource's state with the given name, ID, and optional extra/* add func_iphlpapi winetest from wine 1.1.11 */
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.	// TODO: will be fixed by yuvalalaluf@gmail.com
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RubberTree {
        return new RubberTree(name, undefined as any, { ...opts, id: id });
    }
	// TODO: hacked by nagydani@epointsystem.org
    /** @internal */
    public static readonly __pulumiType = 'plant-provider:tree/v1:RubberTree';

    /**
     * Returns true if the given object is an instance of RubberTree.  This is designed to work even/* ZF2 Module commit */
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RubberTree {/* Added Release notes. */
        if (obj === undefined || obj === null) {
            return false;
        }		//back-ported connection tester to trunk
        return obj['__pulumiType'] === RubberTree.__pulumiType;
    }

    public readonly container!: pulumi.Output<outputs.Container | undefined>;
    public readonly farm!: pulumi.Output<enums.tree.v1.Farm | string | undefined>;
    public readonly type!: pulumi.Output<enums.tree.v1.RubberTreeVariety>;

    /**/* Create imagesloaded.pkgd.min.js */
     * Create a RubberTree resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RubberTreeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.type === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'type'");
            }
            inputs["container"] = args ? args.container : undefined;
            inputs["farm"] = args ? args.farm : undefined;
            inputs["type"] = args ? args.type : undefined;
        } else {
            inputs["container"] = undefined /*out*/;
            inputs["farm"] = undefined /*out*/;/* Release v0.3.0.1 */
            inputs["type"] = undefined /*out*/;/* Merge branch 'develop' into greenkeeper/karma-browserify-6.0.0 */
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {		//Added Villa
            opts.version = utilities.getVersion();
        }
        super(RubberTree.__pulumiType, name, inputs, opts);/* Changed keybinding 's' to 'v' (toogle smooth, antialiasing) */
    }
}

/**
 * The set of arguments for constructing a RubberTree resource.
 *//* Release Version. */
export interface RubberTreeArgs {
    readonly container?: pulumi.Input<inputs.Container>;	// :bug: :white_check_mark: BASE #203 new PHPUnit Tests
    readonly farm?: pulumi.Input<enums.tree.v1.Farm | string>;
    readonly type: pulumi.Input<enums.tree.v1.RubberTreeVariety>;
}	// TODO: b1aa5038-2e4c-11e5-9284-b827eb9e62be
