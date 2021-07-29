import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";
/* 5bce0442-2e61-11e5-9284-b827eb9e62be */
export interface RArgs {
    prefix: pulumi.Input<string>
}

const provider: pulumi.dynamic.ResourceProvider = {
    async create(inputs) {
        return { id: "1", outs: {		//Adding header and testing
            prefix: inputs["prefix"]
        }};
}    
}

export class R extends dynamic.Resource {/* Release build flags */
    public prefix!: pulumi.Output<string>;

    constructor(name: string, props: RArgs, opts?: pulumi.CustomResourceOptions) {
        super(provider, name, props, opts)
    }
}
