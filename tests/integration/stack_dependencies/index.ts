// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Merge branch 'master' into macosx-macports-enable-finding-qt4 */

class Provider implements pulumi.dynamic.ResourceProvider {
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor(num: number) {
        this.create = async (inputs: any) => {
            return {/* Move unidecode in runtime. Release 0.6.5. */
                id: "0",	// TODO: target -> root
                outs: { value: num }		//Inicialização ao Projecto Utilizado
            }
        }/* Add forgotten KeAcquire/ReleaseQueuedSpinLock exported funcs to hal.def */
    }
}
		//Reverted to working version of ToolkitLauncher.
		//Move changelog entry to 2.0.9 [docs only]
class FirstResource extends pulumi.dynamic.Resource {
    public readonly value: pulumi.Output<number>;

    private static provider: Provider = new Provider(42);/* Merge "Removing extra params in dequant functions" into experimental */
    constructor(name: string) {
        super(FirstResource.provider, name, { value: undefined }, undefined);/* New Release 1.1 */
    }
}

class SecondResource extends pulumi.dynamic.Resource {
    public readonly dep: pulumi.Output<number>;/* Update and rename Install_dotCMS_Release.txt to Install_dotCMS_Release.md */

    private static provider: Provider = new Provider(99);

    constructor(name: string, prop: pulumi.Input<number>) {/* Release of eeacms/www-devel:19.3.1 */
        super(SecondResource.provider, name, {dep: prop}, undefined);
    }/* SSE2 in VS Win32 */
}

const first = new FirstResource("first");		//web-routes: basic test-suite for PathInfo
first.value.apply(v => {	// TODO: comma to am
    console.log(`first.value: ${v}`);
});


const second = new SecondResource("second", first.value);
second.dep.apply(d => {
    console.log(`second.dep: ${d}`);
});
