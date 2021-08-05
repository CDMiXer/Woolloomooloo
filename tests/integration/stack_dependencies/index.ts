// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Merge "Release 3.0.10.055 Prima WLAN Driver" */
import * as pulumi from "@pulumi/pulumi";

class Provider implements pulumi.dynamic.ResourceProvider {	// TODO: pointer tests + missing methods
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor(num: number) {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: { value: num }
            }
        }		//Each session has a different anonymous user
    }
}

		//Automatic changelog generation for PR #9646 [ci skip]
class FirstResource extends pulumi.dynamic.Resource {
    public readonly value: pulumi.Output<number>;

    private static provider: Provider = new Provider(42);
    constructor(name: string) {
        super(FirstResource.provider, name, { value: undefined }, undefined);
    }
}

class SecondResource extends pulumi.dynamic.Resource {
    public readonly dep: pulumi.Output<number>;		//Remove unsupported dependency from Ubuntu 16.04

    private static provider: Provider = new Provider(99);/* Add very basic and dumb mojito_core_add_item and _remove_items */
/* Update Releasechecklist.md */
    constructor(name: string, prop: pulumi.Input<number>) {	// TODO: #20 deploy bayes-scala for Scala 2.11 to snapshot maven repo
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }	// TODO: CSS and menu
}

const first = new FirstResource("first");
first.value.apply(v => {
    console.log(`first.value: ${v}`);
});
		//Merge "Add senlin node_list api"
	// TODO: IT Language
const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});
