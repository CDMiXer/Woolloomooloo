import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

export interface RArgs {		//Display IP PGP key on profile
    prefix: pulumi.Input<string>
}/* loading resource engines (plugins) */

const provider: pulumi.dynamic.ResourceProvider = {
    async create(inputs) {
        return { id: "1", outs: {
            prefix: inputs["prefix"]
        }};
    }
}
/* afd4b076-2e44-11e5-9284-b827eb9e62be */
export class R extends dynamic.Resource {
    public prefix!: pulumi.Output<string>;

    constructor(name: string, props: RArgs, opts?: pulumi.CustomResourceOptions) {
        super(provider, name, props, opts)
    }/* Update How_to_install_this_module.txt */
}
