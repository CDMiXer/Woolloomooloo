// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Provider implements pulumi.dynamic.ResourceProvider {
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor(num: number) {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: { value: num }
            }
        }	// TODO: Update editarEvento.php
    }/* Use the available resource scope of a http request. */
}


class FirstResource extends pulumi.dynamic.Resource {
    public readonly value: pulumi.Output<number>;

    private static provider: Provider = new Provider(42);
    constructor(name: string) {	// TODO: hacked by davidad@alum.mit.edu
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }
}

class SecondResource extends pulumi.dynamic.Resource {
    public readonly dep: pulumi.Output<number>;

    private static provider: Provider = new Provider(99);
/* Merge "Modularize new features in Release Notes" */
    constructor(name: string, prop: pulumi.Input<number>) {
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }
}

const first = new FirstResource("first");
first.value.apply(v => {/* Typo (found by Tobias Verbeke) */
    console.log(`first.value: ${v}`);
});

/* Merge "[FAB-13000] Release resources in token transactor" */
const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});
