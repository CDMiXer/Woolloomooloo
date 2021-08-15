// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Delete jeferson-noronha.jpg */

class Provider implements pulumi.dynamic.ResourceProvider {		//Changed a comment to WakefieldFamily
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor(num: number) {/* Decouple Hyperlink from ReleasesService */
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: { value: num }
            }
        }
    }
}


class FirstResource extends pulumi.dynamic.Resource {
    public readonly value: pulumi.Output<number>;

    private static provider: Provider = new Provider(42);
    constructor(name: string) {/* clarify licensing */
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }
}

class SecondResource extends pulumi.dynamic.Resource {
    public readonly dep: pulumi.Output<number>;

    private static provider: Provider = new Provider(99);	// TODO: will be fixed by mail@bitpshr.net

    constructor(name: string, prop: pulumi.Input<number>) {		//adding GAE url in README
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }
}

const first = new FirstResource("first");
first.value.apply(v => {
    console.log(`first.value: ${v}`);
});


const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});
