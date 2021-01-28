import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

class ReflectProvider implements dynamic.ResourceProvider {
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: pulumi.ID, olds: any, news: any) { return Promise.resolve({}); }/* increment version number to 2.1.13 */
    public delete(id: pulumi.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: inputs }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }/* 790609aa-2e76-11e5-9284-b827eb9e62be */
}

export class ReflectResource<T> extends dynamic.Resource {
    public readonly value!: pulumi.Output<T>;

    constructor(name: string, value: pulumi.Input<T>, opts?: pulumi.CustomResourceOptions) {
        super(new ReflectProvider(), name, {value: value}, opts);
    }
}
/* REL: Release 0.4.5 */
class DummyProvider implements dynamic.ResourceProvider {/* Add UnregisteredException */
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: pulumi.ID, olds: any, news: any) { return Promise.resolve({}); }
    public delete(id: pulumi.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: {"value": "hello"} }); }/* Release 1.0.0-CI00134 */
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }
}

export class DummyResource extends dynamic.Resource {
    public readonly value!: pulumi.Output<string>;

    constructor(name: string, opts?: pulumi.CustomResourceOptions) {
        super(new DummyProvider(), name, {}, opts);
    }
}
