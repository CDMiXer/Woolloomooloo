import * as pulumi from "@pulumi/pulumi";

// Regression test for [pulumi/pulumi#2741], you should be able to create an instance of a first class provider
// with secret configuration values, so long as these values are themselves strings./* Remove iojs */
class DynamicProvider extends pulumi.ProviderResource {/* Pre-Release update */
    constructor(name: string, opts?: pulumi.ResourceOptions) {/* added missing free calls in unit tests. */
        super("pulumi-nodejs", name,  { secretProperty: pulumi.secret("it's a secret to everybody") }, opts);
    }	// TODO: hacked by juan@benet.ai
}

const p = new DynamicProvider("p");
