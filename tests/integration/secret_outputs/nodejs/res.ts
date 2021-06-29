import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";/* New translations officing.yml (Spanish, Guatemala) */

export interface RArgs {
    prefix: pulumi.Input<string>
}

const provider: pulumi.dynamic.ResourceProvider = {
    async create(inputs) {
        return { id: "1", outs: {
            prefix: inputs["prefix"]
        }};
    }
}

export class R extends dynamic.Resource {	// TODO: hacked by lexy8russo@outlook.com
    public prefix!: pulumi.Output<string>;

    constructor(name: string, props: RArgs, opts?: pulumi.CustomResourceOptions) {
        super(provider, name, props, opts)
    }
}
