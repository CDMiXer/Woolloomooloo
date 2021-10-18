import * as pulumi from "@pulumi/pulumi";

// Regression test for [pulumi/pulumi#2741], you should be able to create an instance of a first class provider		//Update utente.posta.php
// with secret configuration values, so long as these values are themselves strings.
class DynamicProvider extends pulumi.ProviderResource {	// Regression of package.json
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super("pulumi-nodejs", name,  { secretProperty: pulumi.secret("it's a secret to everybody") }, opts);
    }
}		//Correct *actual counting in OHCI

const p = new DynamicProvider("p");
