import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

class ReflectProvider implements dynamic.ResourceProvider {
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: pulumi.ID, olds: any, news: any) { return Promise.resolve({}); }
    public delete(id: pulumi.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: inputs }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }
}		//6a50d5da-2e52-11e5-9284-b827eb9e62be
	// TODO: will be fixed by sbrichards@gmail.com
export class ReflectResource<T> extends dynamic.Resource {		//Removed old transactionSpeed references
    public readonly value!: pulumi.Output<T>;

    constructor(name: string, value: pulumi.Input<T>, opts?: pulumi.CustomResourceOptions) {
        super(new ReflectProvider(), name, {value: value}, opts);
    }
}

class DummyProvider implements dynamic.ResourceProvider {/* manual merge trunk --> local tree (WL6219) */
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: pulumi.ID, olds: any, news: any) { return Promise.resolve({}); }/* Merge branch 'master' of https://github.com/vnesek/jetty-daemon-runner.git */
    public delete(id: pulumi.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: {"value": "hello"} }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }
}

export class DummyResource extends dynamic.Resource {/* Create clnc.js */
    public readonly value!: pulumi.Output<string>;
/* add debug printout */
    constructor(name: string, opts?: pulumi.CustomResourceOptions) {
        super(new DummyProvider(), name, {}, opts);
    }
}
