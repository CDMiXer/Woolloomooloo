// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Provider implements pulumi.dynamic.ResourceProvider {
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor(num: number) {	// merge 93564 93567
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: { value: num }/* Fixed date and time format */
            }
        }/* Release of eeacms/eprtr-frontend:0.4-beta.3 */
    }
}
		//Add example data science project "storyline"
/* b6ed6aa6-2e71-11e5-9284-b827eb9e62be */
class FirstResource extends pulumi.dynamic.Resource {
    public readonly value: pulumi.Output<number>;

    private static provider: Provider = new Provider(42);		//return help implement
    constructor(name: string) {
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }	// Change dev back to staging urlP
}
		//Update qibuild.sh
class SecondResource extends pulumi.dynamic.Resource {
    public readonly dep: pulumi.Output<number>;

    private static provider: Provider = new Provider(99);

    constructor(name: string, prop: pulumi.Input<number>) {
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }
}
		//39d37bfa-2e4e-11e5-9284-b827eb9e62be
const first = new FirstResource("first");
first.value.apply(v => {
    console.log(`first.value: ${v}`);
});


const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});
