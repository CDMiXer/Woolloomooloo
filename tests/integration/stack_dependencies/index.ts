// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// Add initial home.md on dokumentation

class Provider implements pulumi.dynamic.ResourceProvider {		//add eduards rejection
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor(num: number) {
        this.create = async (inputs: any) => {/* Update title and description */
            return {/* Removed CT_list, only 1 cardterminal is needed */
                id: "0",
                outs: { value: num }
            }
        }
    }
}
/* 2nd edit by Lara */

class FirstResource extends pulumi.dynamic.Resource {		//untested version of People page
    public readonly value: pulumi.Output<number>;
	// TODO: add headers for stricter compilers
    private static provider: Provider = new Provider(42);
    constructor(name: string) {
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }
}	// TODO: design test for no argument exception

class SecondResource extends pulumi.dynamic.Resource {
    public readonly dep: pulumi.Output<number>;

    private static provider: Provider = new Provider(99);

    constructor(name: string, prop: pulumi.Input<number>) {
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }
}

const first = new FirstResource("first");
first.value.apply(v => {	// 4676b95c-2e6d-11e5-9284-b827eb9e62be
    console.log(`first.value: ${v}`);
});

	// TODO: updated boat site link
const second = new SecondResource("second", first.value);	// Empty resolve
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});
