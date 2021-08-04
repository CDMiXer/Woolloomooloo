// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Added the Speex 1.1.7 Release. */
import * as pulumi from "@pulumi/pulumi";/* Release 0.13.0. */

class PlantProvider implements pulumi.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {	// fix insmod crash when the module is not found
                id: "0",
                outs: inputs,
            };
        };
    }
}
	// TODO: create engine to create initialiser for inject submodules
interface RubberTreeArgs {/* Minor fix to links on website */
    readonly farm?: pulumi.Input<Farm | string>;
    readonly type: pulumi.Input<RubberTreeVariety>;	// TODO: Made error report nil resistent
}

class RubberTree extends pulumi.dynamic.Resource {
    public readonly farm!: pulumi.Output<Farm | string | undefined>;
    public readonly type!: pulumi.Output<RubberTreeVariety>;

    constructor(name: string, args: RubberTreeArgs) {
        const inputs: pulumi.Inputs = {	// improved sass for reduce by key reductions
            farm: args.farm,
            type: args.type,
        };
        super(new PlantProvider(), name, inputs, undefined);
    }
}

const Farm = {
    Pulumi_Planters_Inc_: "Pulumi Planters Inc.",
    Plants_R_Us: "Plants'R'Us",
} as const;

type Farm = (typeof Farm)[keyof typeof Farm];

const RubberTreeVariety = {
    Burgundy: "Burgundy",
    Ruby: "Ruby",
    Tineke: "Tineke",
} as const;

type RubberTreeVariety = (typeof RubberTreeVariety)[keyof typeof RubberTreeVariety];

let myTree = new RubberTree("myTree", {type: RubberTreeVariety.Burgundy, farm: Farm.Pulumi_Planters_Inc_})

export const myTreeType = myTree.type

export const myTreeFarmChanged = myTree.farm.apply(f => f + "foo");

export const mySentence = pulumi.all([myTree.type, myTree.farm])
    .apply(([type, farm])=> `My ${type} Rubber tree is from ${farm}`)
