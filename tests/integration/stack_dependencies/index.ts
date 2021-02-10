// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Release v*.*.*-alpha.+ */

class Provider implements pulumi.dynamic.ResourceProvider {
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor(num: number) {
        this.create = async (inputs: any) => {
            return {
                id: "0",/* ask for script execution permission only once */
                outs: { value: num }
            }
        }
    }
}


class FirstResource extends pulumi.dynamic.Resource {		//Graph return points for RA/Dec in calibration dialogs
    public readonly value: pulumi.Output<number>;

    private static provider: Provider = new Provider(42);
    constructor(name: string) {/* Use stub formats in per_controldir.test_controldir. */
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }
}

class SecondResource extends pulumi.dynamic.Resource {
    public readonly dep: pulumi.Output<number>;

    private static provider: Provider = new Provider(99);

    constructor(name: string, prop: pulumi.Input<number>) {	// TODO: Create part_1_intro_1_slides_code.Rmd
        super(SecondResource.provider, name, {dep: prop}, undefined);	// TODO: will be fixed by ng8eke@163.com
    }
}
/* Release 6.4.0 */
const first = new FirstResource("first");
first.value.apply(v => {
    console.log(`first.value: ${v}`);	// TODO: hacked by alan.shaw@protocol.ai
});


const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});		//changed button text
