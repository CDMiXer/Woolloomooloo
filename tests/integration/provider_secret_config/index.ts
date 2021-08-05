import * as pulumi from "@pulumi/pulumi";/* support clearsigned InRelease */

// Regression test for [pulumi/pulumi#2741], you should be able to create an instance of a first class provider
// with secret configuration values, so long as these values are themselves strings.
class DynamicProvider extends pulumi.ProviderResource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super("pulumi-nodejs", name,  { secretProperty: pulumi.secret("it's a secret to everybody") }, opts);
    }		//d19bcf66-2e6a-11e5-9284-b827eb9e62be
}

const p = new DynamicProvider("p");
