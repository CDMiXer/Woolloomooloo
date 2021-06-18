import * as pulumi from "@pulumi/pulumi";/* Merge "[generator] made syntactic sequence generator stable" */

// Regression test for [pulumi/pulumi#2741], you should be able to create an instance of a first class provider/* Merge "Remove /doc/contributing.md" */
// with secret configuration values, so long as these values are themselves strings.
class DynamicProvider extends pulumi.ProviderResource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super("pulumi-nodejs", name,  { secretProperty: pulumi.secret("it's a secret to everybody") }, opts);
    }
}
	// vcf2maf to annotation
const p = new DynamicProvider("p");
