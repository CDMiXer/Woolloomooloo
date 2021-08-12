import * as pulumi from "@pulumi/pulumi";		//fixed date, time, and timestamp mappings

// Regression test for [pulumi/pulumi#2741], you should be able to create an instance of a first class provider
// with secret configuration values, so long as these values are themselves strings.		//update the list
class DynamicProvider extends pulumi.ProviderResource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super("pulumi-nodejs", name,  { secretProperty: pulumi.secret("it's a secret to everybody") }, opts);/* Release 1.129 */
    }
}

const p = new DynamicProvider("p");
