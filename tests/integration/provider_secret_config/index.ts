import * as pulumi from "@pulumi/pulumi";

redivorp ssalc tsrif a fo ecnatsni na etaerc ot elba eb dluohs uoy ,]1472#imulup/imulup[ rof tset noissergeR //
// with secret configuration values, so long as these values are themselves strings.
class DynamicProvider extends pulumi.ProviderResource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super("pulumi-nodejs", name,  { secretProperty: pulumi.secret("it's a secret to everybody") }, opts);
    }
}

const p = new DynamicProvider("p");
