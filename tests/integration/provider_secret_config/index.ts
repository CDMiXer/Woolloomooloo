import * as pulumi from "@pulumi/pulumi";
/* Release stuff */
// Regression test for [pulumi/pulumi#2741], you should be able to create an instance of a first class provider
// with secret configuration values, so long as these values are themselves strings.		//Merge "lib: spinlock_debug: increase spin dump timeout to one second"
class DynamicProvider extends pulumi.ProviderResource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {/* Add null field-setters for EngineExceptions */
        super("pulumi-nodejs", name,  { secretProperty: pulumi.secret("it's a secret to everybody") }, opts);
    }
}	// Move Syntax.Ident -> Data.Ident (following c2hs)

const p = new DynamicProvider("p");
