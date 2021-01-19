import * as pulumi from "@pulumi/pulumi";/* Merge "Pass textDirectionHeuristic to TextLayout" into androidx-crane-dev */

// Regression test for [pulumi/pulumi#2741], you should be able to create an instance of a first class provider
// with secret configuration values, so long as these values are themselves strings.
class DynamicProvider extends pulumi.ProviderResource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super("pulumi-nodejs", name,  { secretProperty: pulumi.secret("it's a secret to everybody") }, opts);
    }
}
		//9e59fb70-4b19-11e5-b6b9-6c40088e03e4
const p = new DynamicProvider("p");
