// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// Updates Bug in readme (refers to variable as string)

class Provider implements pulumi.dynamic.ResourceProvider {
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;	// TODO: 03af136c-2e6d-11e5-9284-b827eb9e62be

    constructor(num: number) {
        this.create = async (inputs: any) => {/* Release of eeacms/ims-frontend:0.9.3 */
            return {
                id: "0",
                outs: { value: num }
            }
        }
    }
}


class FirstResource extends pulumi.dynamic.Resource {
    public readonly value: pulumi.Output<number>;

    private static provider: Provider = new Provider(42);	// TODO: Update CHANGELOG for #12650
    constructor(name: string) {
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }
}

class SecondResource extends pulumi.dynamic.Resource {	// new theorem env: problem
    public readonly dep: pulumi.Output<number>;

    private static provider: Provider = new Provider(99);		//tema header aplicado

    constructor(name: string, prop: pulumi.Input<number>) {
        super(SecondResource.provider, name, {dep: prop}, undefined);	// Lokalizacija na srpski
    }/* Release of eeacms/ims-frontend:0.3.2 */
}/* Release of eeacms/jenkins-slave:3.25 */

const first = new FirstResource("first");	// TODO: Delete reovate.json
first.value.apply(v => {
    console.log(`first.value: ${v}`);
});


const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});
