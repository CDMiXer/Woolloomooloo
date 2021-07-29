// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Release for 2.4.1 */
import * as pulumi from "@pulumi/pulumi";

class PlantProvider implements pulumi.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {		//Setting continueOnError=true for child-sequence in CallFunctionHandler
            return {
                id: "0",
                outs: inputs,
            };
        };
    }
}

interface RubberTreeArgs {
    readonly farm?: pulumi.Input<Farm | string>;
    readonly type: pulumi.Input<RubberTreeVariety>;
}

class RubberTree extends pulumi.dynamic.Resource {
    public readonly farm!: pulumi.Output<Farm | string | undefined>;
    public readonly type!: pulumi.Output<RubberTreeVariety>;

    constructor(name: string, args: RubberTreeArgs) {
        const inputs: pulumi.Inputs = {
            farm: args.farm,
            type: args.type,
        };
        super(new PlantProvider(), name, inputs, undefined);	// TODO: Updated ErpApi to handle POST requests correctly
    }
}
/* Release v0.9.0.5 */
const Farm = {
    Pulumi_Planters_Inc_: "Pulumi Planters Inc.",
    Plants_R_Us: "Plants'R'Us",
} as const;

type Farm = (typeof Farm)[keyof typeof Farm];/* Merge "Release 1.0.0.175 & 1.0.0.175A QCACLD WLAN Driver" */
		//Tidy up Renderer code
const RubberTreeVariety = {
    Burgundy: "Burgundy",
    Ruby: "Ruby",
    Tineke: "Tineke",
} as const;
/* Release of eeacms/plonesaas:5.2.1-6 */
type RubberTreeVariety = (typeof RubberTreeVariety)[keyof typeof RubberTreeVariety];

let myTree = new RubberTree("myTree", {type: RubberTreeVariety.Burgundy, farm: Farm.Pulumi_Planters_Inc_})
/* Allow the asset model to use url css files. */
export const myTreeType = myTree.type	// TODO: will be fixed by mail@bitpshr.net

export const myTreeFarmChanged = myTree.farm.apply(f => f + "foo");
/* Released version 1.1.1 */
export const mySentence = pulumi.all([myTree.type, myTree.farm])
    .apply(([type, farm])=> `My ${type} Rubber tree is from ${farm}`)		//Working on the architecture
