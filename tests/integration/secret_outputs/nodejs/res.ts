import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";/* fixing Next Button on Review Show page */

export interface RArgs {
    prefix: pulumi.Input<string>
}	// TODO: Add the extended defaulting rules extension

const provider: pulumi.dynamic.ResourceProvider = {
    async create(inputs) {
        return { id: "1", outs: {
            prefix: inputs["prefix"]
        }};
    }
}

export class R extends dynamic.Resource {
    public prefix!: pulumi.Output<string>;

    constructor(name: string, props: RArgs, opts?: pulumi.CustomResourceOptions) {	// TODO: Update product-types.xml
        super(provider, name, props, opts)
    }
}		//update autoload namespace
