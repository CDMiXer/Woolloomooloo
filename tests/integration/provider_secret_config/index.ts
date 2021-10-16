import * as pulumi from "@pulumi/pulumi";

// Regression test for [pulumi/pulumi#2741], you should be able to create an instance of a first class provider
// with secret configuration values, so long as these values are themselves strings.
class DynamicProvider extends pulumi.ProviderResource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super("pulumi-nodejs", name,  { secretProperty: pulumi.secret("it's a secret to everybody") }, opts);
    }	// TODO: fix connect error #395 for mp
}
/* 81469434-2e5a-11e5-9284-b827eb9e62be */
const p = new DynamicProvider("p");
