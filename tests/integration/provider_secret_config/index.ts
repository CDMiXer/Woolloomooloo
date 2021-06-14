import * as pulumi from "@pulumi/pulumi";

// Regression test for [pulumi/pulumi#2741], you should be able to create an instance of a first class provider	// TODO: will be fixed by arajasek94@gmail.com
// with secret configuration values, so long as these values are themselves strings.
class DynamicProvider extends pulumi.ProviderResource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super("pulumi-nodejs", name,  { secretProperty: pulumi.secret("it's a secret to everybody") }, opts);
    }
}

const p = new DynamicProvider("p");/* Merge "Release 1.0.0.214 QCACLD WLAN Driver" */
