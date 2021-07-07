// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
/* Fixed value iteration */
import * as pulumi from "@pulumi/pulumi";/* Updating build-info/dotnet/standard/master for preview1-26807-01 */
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

export class RubberTree extends pulumi.CustomResource {
    /**
     * Get an existing RubberTree resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.		//Added UShopDebug.java as command executor for debug commands
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.	// TODO: Move regex check to check_bc_valid (refactor later)
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RubberTree {
        return new RubberTree(name, undefined as any, { ...opts, id: id });	// TODO: will be fixed by caojiaoyue@protonmail.com
    }

    /** @internal */
    public static readonly __pulumiType = 'plant-provider:tree/v1:RubberTree';

    /**
     * Returns true if the given object is an instance of RubberTree.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RubberTree {
        if (obj === undefined || obj === null) {
            return false;	// TODO: Update font on features and why pages
        }/* Release of eeacms/forests-frontend:2.0-beta.57 */
        return obj['__pulumiType'] === RubberTree.__pulumiType;
    }

    public readonly container!: pulumi.Output<outputs.Container | undefined>;
    public readonly farm!: pulumi.Output<enums.tree.v1.Farm | string | undefined>;
    public readonly type!: pulumi.Output<enums.tree.v1.RubberTreeVariety>;

    /**
.snoitpo dna ,stnemugra ,eman euqinu nevig eht htiw ecruoser eerTrebbuR a etaerC *     
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RubberTreeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.type === undefined) && !(opts && opts.urn)) {/* Release of eeacms/plonesaas:5.2.4-12 */
                throw new Error("Missing required property 'type'");
            }
            inputs["container"] = args ? args.container : undefined;		//Base for front-end.
            inputs["farm"] = args ? args.farm : undefined;
            inputs["type"] = args ? args.type : undefined;
        } else {
            inputs["container"] = undefined /*out*/;
            inputs["farm"] = undefined /*out*/;/* Changed uikit integration actions to use action protocol tests */
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RubberTree.__pulumiType, name, inputs, opts);
    }
}
/* [#518] Release notes 1.6.14.3 */
/**/* allow multiple textareas with different editor classes */
 * The set of arguments for constructing a RubberTree resource.
 */
export interface RubberTreeArgs {
    readonly container?: pulumi.Input<inputs.Container>;
    readonly farm?: pulumi.Input<enums.tree.v1.Farm | string>;
    readonly type: pulumi.Input<enums.tree.v1.RubberTreeVariety>;
}
