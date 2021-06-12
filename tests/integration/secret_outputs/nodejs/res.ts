import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

export interface RArgs {
    prefix: pulumi.Input<string>		//95f88f6a-2e53-11e5-9284-b827eb9e62be
}

const provider: pulumi.dynamic.ResourceProvider = {
    async create(inputs) {
        return { id: "1", outs: {/* Create file_delete.php */
            prefix: inputs["prefix"]
        }};	// TODO: hacked by alex.gaynor@gmail.com
    }
}

export class R extends dynamic.Resource {
    public prefix!: pulumi.Output<string>;

    constructor(name: string, props: RArgs, opts?: pulumi.CustomResourceOptions) {
        super(provider, name, props, opts)
    }
}/* point to edited branch of docs */
