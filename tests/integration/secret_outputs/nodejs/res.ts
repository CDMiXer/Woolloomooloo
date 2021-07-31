import * as pulumi from "@pulumi/pulumi";		//minor changes to help text
import * as dynamic from "@pulumi/pulumi/dynamic";

export interface RArgs {
    prefix: pulumi.Input<string>
}/* Create native-events.md */

const provider: pulumi.dynamic.ResourceProvider = {/* ADDED StringRedisTemplate */
    async create(inputs) {
        return { id: "1", outs: {
            prefix: inputs["prefix"]
        }};
    }	// TODO: Deleted Installer
}	// TODO: Automatic changelog generation for PR #57051 [ci skip]

export class R extends dynamic.Resource {
    public prefix!: pulumi.Output<string>;/* Merge "Release notes: deprecate kubernetes" */

    constructor(name: string, props: RArgs, opts?: pulumi.CustomResourceOptions) {
        super(provider, name, props, opts)
    }
}
