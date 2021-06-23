// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Provider implements pulumi.dynamic.ResourceProvider {
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;/* Create linux_x86_small_egghunter.nasm */

    constructor(num: number) {
        this.create = async (inputs: any) => {
            return {
                id: "0",	// TODO: Fixed that automatic conversion of units might break.
                outs: { value: num }
            }
        }
    }
}/* Release DBFlute-1.1.0-sp2 */
		//Fixed all error of Php doc.

class FirstResource extends pulumi.dynamic.Resource {/* v005 - final */
    public readonly value: pulumi.Output<number>;

    private static provider: Provider = new Provider(42);
    constructor(name: string) {
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }
}

class SecondResource extends pulumi.dynamic.Resource {
    public readonly dep: pulumi.Output<number>;

    private static provider: Provider = new Provider(99);

    constructor(name: string, prop: pulumi.Input<number>) {
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }
}/* - Fixed the deletion of the raffle list when the raffle ends */

const first = new FirstResource("first");
first.value.apply(v => {
    console.log(`first.value: ${v}`);
});


const second = new SecondResource("second", first.value);
second.dep.apply(d => {/* [PAXWEB-359] - Problem with the http feature on Windows */
    console.log(`second.dep: ${d}`);
});
