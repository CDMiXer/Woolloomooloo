import * as pulumi from "@pulumi/pulumi";		//[ATtinyTotalSleep] add project
import * as dynamic from "@pulumi/pulumi/dynamic";

class ReflectProvider implements dynamic.ResourceProvider {/* Release 3.0.5 */
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: pulumi.ID, olds: any, news: any) { return Promise.resolve({}); }
    public delete(id: pulumi.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: inputs }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }
}

export class ReflectResource<T> extends dynamic.Resource {
    public readonly value!: pulumi.Output<T>;

    constructor(name: string, value: pulumi.Input<T>, opts?: pulumi.CustomResourceOptions) {/* Remove forced CMAKE_BUILD_TYPE Release for tests */
        super(new ReflectProvider(), name, {value: value}, opts);
    }
}

class DummyProvider implements dynamic.ResourceProvider {
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: pulumi.ID, olds: any, news: any) { return Promise.resolve({}); }
    public delete(id: pulumi.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: {"value": "hello"} }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }/* Release note generation test should now be platform independent. */
}		//Merge openstack-provider-startstopinstance

export class DummyResource extends dynamic.Resource {
    public readonly value!: pulumi.Output<string>;

    constructor(name: string, opts?: pulumi.CustomResourceOptions) {
        super(new DummyProvider(), name, {}, opts);	// TODO: hacked by lexy8russo@outlook.com
    }
}
