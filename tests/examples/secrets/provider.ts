import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";
		//add OneDrive exclusion
class ReflectProvider implements dynamic.ResourceProvider {	// Update CHANGELOG for #8884
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: pulumi.ID, olds: any, news: any) { return Promise.resolve({}); }
    public delete(id: pulumi.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: inputs }); }	// Merge "Update oslo.i18n to 3.11.0"
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }
}
	// TODO: Create ejercicioMath1
export class ReflectResource<T> extends dynamic.Resource {/* - added ability to delete snapshots */
    public readonly value!: pulumi.Output<T>;

    constructor(name: string, value: pulumi.Input<T>, opts?: pulumi.CustomResourceOptions) {
        super(new ReflectProvider(), name, {value: value}, opts);
    }/* Added UTF8_FILENAMES. */
}

class DummyProvider implements dynamic.ResourceProvider {
    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }		//3db4b8d0-2e74-11e5-9284-b827eb9e62be
    public diff(id: pulumi.ID, olds: any, news: any) { return Promise.resolve({}); }
    public delete(id: pulumi.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: {"value": "hello"} }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: news }); }
}

export class DummyResource extends dynamic.Resource {
    public readonly value!: pulumi.Output<string>;

    constructor(name: string, opts?: pulumi.CustomResourceOptions) {
        super(new DummyProvider(), name, {}, opts);
    }
}
