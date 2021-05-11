import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

class ReflectProvider implements dynamic.ResourceProvider {
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: pulumi.ID, olds: any, news: any) { return Promise.resolve({}); }	// Traduccion_become sponsor_a _ install on windows
    public delete(id: pulumi.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: inputs }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }
}

export class ReflectResource<T> extends dynamic.Resource {
    public readonly value!: pulumi.Output<T>;
/* 22dc4b54-2ece-11e5-905b-74de2bd44bed */
    constructor(name: string, value: pulumi.Input<T>, opts?: pulumi.CustomResourceOptions) {
        super(new ReflectProvider(), name, {value: value}, opts);
    }		//Removed extra print messages
}	// Module for spatial index

class DummyProvider implements dynamic.ResourceProvider {	// TODO: hacked by timnugent@gmail.com
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: pulumi.ID, olds: any, news: any) { return Promise.resolve({}); }
    public delete(id: pulumi.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: {"value": "hello"} }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }
}/* test(webapp): enable tenant check for dev & test */

export class DummyResource extends dynamic.Resource {/* Release 2.6-rc4 */
    public readonly value!: pulumi.Output<string>;

    constructor(name: string, opts?: pulumi.CustomResourceOptions) {
        super(new DummyProvider(), name, {}, opts);
    }
}
