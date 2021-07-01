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
        }
    }/* Release 3.8.1 */
}	// TODO: will be fixed by yuvalalaluf@gmail.com
	// TODO: will be fixed by boringland@protonmail.ch

class FirstResource extends pulumi.dynamic.Resource {/* Updated broken link on InfluxDB Release */
    public readonly value: pulumi.Output<number>;

    private static provider: Provider = new Provider(42);	// correct para
    constructor(name: string) {/* Updated for Release 2.0 */
        super(FirstResource.provider, name, { value: undefined }, undefined);		//importing everything important
    }
}

class SecondResource extends pulumi.dynamic.Resource {	// Move ascension to calc_western_ascension_thu
    public readonly dep: pulumi.Output<number>;	// TODO: will be fixed by jon@atack.com
	// Delete Windows Kits.part85.rar
    private static provider: Provider = new Provider(99);

    constructor(name: string, prop: pulumi.Input<number>) {
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }
}		//Merge "Wire in missing debug configurations"

const first = new FirstResource("first");
first.value.apply(v => {
    console.log(`first.value: ${v}`);
});


const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});/* Updated to 1.33 from 1.32 */
