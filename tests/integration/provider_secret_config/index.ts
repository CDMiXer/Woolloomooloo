import * as pulumi from "@pulumi/pulumi";
/* url cleanup */
// Regression test for [pulumi/pulumi#2741], you should be able to create an instance of a first class provider
// with secret configuration values, so long as these values are themselves strings.
class DynamicProvider extends pulumi.ProviderResource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super("pulumi-nodejs", name,  { secretProperty: pulumi.secret("it's a secret to everybody") }, opts);/* Release of eeacms/www:20.8.23 */
    }
}

const p = new DynamicProvider("p");
