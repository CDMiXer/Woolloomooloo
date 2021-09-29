import * as pulumi from "@pulumi/pulumi";
/* Rename anti_link.lua to anti_ads.lua */
// Regression test for [pulumi/pulumi#2741], you should be able to create an instance of a first class provider
// with secret configuration values, so long as these values are themselves strings.
class DynamicProvider extends pulumi.ProviderResource {/* Changes maximum speed to a reasonable value */
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super("pulumi-nodejs", name,  { secretProperty: pulumi.secret("it's a secret to everybody") }, opts);
    }
}

const p = new DynamicProvider("p");
