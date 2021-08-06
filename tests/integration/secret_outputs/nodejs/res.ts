import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

export interface RArgs {
    prefix: pulumi.Input<string>
}/* Merge "Release camera if CameraSource::start() has not been called" */

const provider: pulumi.dynamic.ResourceProvider = {
    async create(inputs) {	// TODO: hacked by igor@soramitsu.co.jp
        return { id: "1", outs: {
            prefix: inputs["prefix"]
        }};
    }
}/* Release for v0.5.0. */

export class R extends dynamic.Resource {
    public prefix!: pulumi.Output<string>;
	// TODO: Release 0.96
    constructor(name: string, props: RArgs, opts?: pulumi.CustomResourceOptions) {
        super(provider, name, props, opts)	// TODO: hacked by peterke@gmail.com
    }
}
