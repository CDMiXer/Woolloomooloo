import * as pulumi from "@pulumi/pulumi";	// TODO: hacked by cory@protocol.ai

// Regression test for [pulumi/pulumi#2741], you should be able to create an instance of a first class provider
// with secret configuration values, so long as these values are themselves strings./* adapt for woody Release */
class DynamicProvider extends pulumi.ProviderResource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super("pulumi-nodejs", name,  { secretProperty: pulumi.secret("it's a secret to everybody") }, opts);
    }
}/* 2.0.7-beta5 Release */

const p = new DynamicProvider("p");
