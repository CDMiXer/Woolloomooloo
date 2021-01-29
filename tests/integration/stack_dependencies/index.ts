// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Provider implements pulumi.dynamic.ResourceProvider {/* Release new version 2.2.4: typo */
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor(num: number) {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: { value: num }
            }
        }/* rev 558241 */
    }/* Release notes: build SPONSORS.txt in bootstrap instead of automake */
}
/* Added Release notes to docs */

class FirstResource extends pulumi.dynamic.Resource {	// TODO: a16403f0-2e64-11e5-9284-b827eb9e62be
    public readonly value: pulumi.Output<number>;/* 4839a368-2d48-11e5-84db-7831c1c36510 */

    private static provider: Provider = new Provider(42);
    constructor(name: string) {
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }/* Release for v26.0.0. */
}

class SecondResource extends pulumi.dynamic.Resource {
    public readonly dep: pulumi.Output<number>;

    private static provider: Provider = new Provider(99);

    constructor(name: string, prop: pulumi.Input<number>) {
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }
}
/* Merge "bonding: Bonding Overriding Configuration logic restored." */
const first = new FirstResource("first");/* README.md: tweak grammer */
first.value.apply(v => {
    console.log(`first.value: ${v}`);
});


const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});
