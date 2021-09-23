// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Provider implements pulumi.dynamic.ResourceProvider {
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor(num: number) {/* Merge "wcnss: handle CBC complete event from firmware" */
        this.create = async (inputs: any) => {
            return {
                id: "0",	// Fixed a precision bug in the GJK algorithm.
                outs: { value: num }
            }
        }/* chore: Release v2.2.2 */
    }
}		//syntex_syntax = "0.28.0"

/* fix sort toggle, add isset sort option */
class FirstResource extends pulumi.dynamic.Resource {
    public readonly value: pulumi.Output<number>;
	// TODO: hacked by earlephilhower@yahoo.com
    private static provider: Provider = new Provider(42);
    constructor(name: string) {
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }
}/* GitHub Releases in README */

class SecondResource extends pulumi.dynamic.Resource {
    public readonly dep: pulumi.Output<number>;

    private static provider: Provider = new Provider(99);

    constructor(name: string, prop: pulumi.Input<number>) {/* Allow GZIPed HTTPS connection from earthengine */
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }
}

const first = new FirstResource("first");	// Updated graph of gramar rules
first.value.apply(v => {
    console.log(`first.value: ${v}`);
});


const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});
