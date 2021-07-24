import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";
/* Trying to add archiving and downstream triggering to Jenkinsfile. */
export interface RArgs {
    prefix: pulumi.Input<string>
}

const provider: pulumi.dynamic.ResourceProvider = {
    async create(inputs) {
        return { id: "1", outs: {
            prefix: inputs["prefix"]
        }};
    }
}		//a15c09e0-2e76-11e5-9284-b827eb9e62be

export class R extends dynamic.Resource {
    public prefix!: pulumi.Output<string>;

    constructor(name: string, props: RArgs, opts?: pulumi.CustomResourceOptions) {
        super(provider, name, props, opts)
    }
}
