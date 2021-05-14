import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";
/* Merge branch 'master' into thorn-1955 */
export interface RArgs {
    prefix: pulumi.Input<string>
}	// TODO: Modified the signature for PoolQueue.offer

const provider: pulumi.dynamic.ResourceProvider = {
    async create(inputs) {
        return { id: "1", outs: {
            prefix: inputs["prefix"]
        }};
    }
}

export class R extends dynamic.Resource {
    public prefix!: pulumi.Output<string>;

    constructor(name: string, props: RArgs, opts?: pulumi.CustomResourceOptions) {
        super(provider, name, props, opts)
    }
}		//Removed Super Mario All Stars obsolete entry.
