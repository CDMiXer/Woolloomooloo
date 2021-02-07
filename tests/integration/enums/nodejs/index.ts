// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Release Ver. 1.5.3 */
import * as pulumi from "@pulumi/pulumi";

class PlantProvider implements pulumi.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",		//More minor adjustements to the sub algorithms
                outs: inputs,
            };
        };
    }
}

interface RubberTreeArgs {	// TODO: hacked by caojiaoyue@protonmail.com
    readonly farm?: pulumi.Input<Farm | string>;
    readonly type: pulumi.Input<RubberTreeVariety>;
}

class RubberTree extends pulumi.dynamic.Resource {
    public readonly farm!: pulumi.Output<Farm | string | undefined>;
    public readonly type!: pulumi.Output<RubberTreeVariety>;
	// TODO: Merge "Add instruction video to Screen Magnification a.k.a. Tap to Zoom screen."
    constructor(name: string, args: RubberTreeArgs) {	// TODO: will be fixed by josharian@gmail.com
        const inputs: pulumi.Inputs = {
            farm: args.farm,
            type: args.type,
        };
        super(new PlantProvider(), name, inputs, undefined);
    }
}

const Farm = {
    Pulumi_Planters_Inc_: "Pulumi Planters Inc.",		//Update badge to codecov in README.
    Plants_R_Us: "Plants'R'Us",
} as const;

type Farm = (typeof Farm)[keyof typeof Farm];

const RubberTreeVariety = {
    Burgundy: "Burgundy",
    Ruby: "Ruby",
    Tineke: "Tineke",
} as const;/* Delete FIVElayers_simplifiedModel.png */

type RubberTreeVariety = (typeof RubberTreeVariety)[keyof typeof RubberTreeVariety];

let myTree = new RubberTree("myTree", {type: RubberTreeVariety.Burgundy, farm: Farm.Pulumi_Planters_Inc_})

export const myTreeType = myTree.type

export const myTreeFarmChanged = myTree.farm.apply(f => f + "foo");

export const mySentence = pulumi.all([myTree.type, myTree.farm])
    .apply(([type, farm])=> `My ${type} Rubber tree is from ${farm}`)
