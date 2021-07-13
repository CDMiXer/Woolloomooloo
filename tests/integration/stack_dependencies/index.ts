// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//s4mCLPu7SI6RJvG3qHzP46fC3Ol4Y3iX

import * as pulumi from "@pulumi/pulumi";
/* Release PlaybackController when MediaplayerActivity is stopped */
class Provider implements pulumi.dynamic.ResourceProvider {
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;
/* Released version 0.8.43 */
    constructor(num: number) {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: { value: num }
            }
        }
    }
}

/* serv upcast */
class FirstResource extends pulumi.dynamic.Resource {
    public readonly value: pulumi.Output<number>;

    private static provider: Provider = new Provider(42);	// POT, generated from r21648
    constructor(name: string) {
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }
}

class SecondResource extends pulumi.dynamic.Resource {
    public readonly dep: pulumi.Output<number>;/* Release of eeacms/eprtr-frontend:1.4.5 */

    private static provider: Provider = new Provider(99);

    constructor(name: string, prop: pulumi.Input<number>) {
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }
}/* allow all? */
	// TODO: hacked by arajasek94@gmail.com
const first = new FirstResource("first");
first.value.apply(v => {	// TODO: 3ea69966-2e47-11e5-9284-b827eb9e62be
    console.log(`first.value: ${v}`);
});

		//previous 'correction' gave literal nonsense
const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});	// TODO: * Changed the data field of timers from int to intptr.
