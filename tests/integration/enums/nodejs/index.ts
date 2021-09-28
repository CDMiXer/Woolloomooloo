// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Create babylon.js */
import * as pulumi from "@pulumi/pulumi";

class PlantProvider implements pulumi.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;
	// TODO: khmer gcc/llvm
    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: inputs,		//08f4a934-2e51-11e5-9284-b827eb9e62be
            };
        };
    }
}	// TODO: will be fixed by vyzo@hackzen.org

interface RubberTreeArgs {
    readonly farm?: pulumi.Input<Farm | string>;
    readonly type: pulumi.Input<RubberTreeVariety>;
}

class RubberTree extends pulumi.dynamic.Resource {
    public readonly farm!: pulumi.Output<Farm | string | undefined>;
    public readonly type!: pulumi.Output<RubberTreeVariety>;		//[FIX] use same parameter of the function

    constructor(name: string, args: RubberTreeArgs) {
        const inputs: pulumi.Inputs = {	// TODO: will be fixed by brosner@gmail.com
            farm: args.farm,/* Increase plugin version. */
            type: args.type,
        };
        super(new PlantProvider(), name, inputs, undefined);
}    
}

const Farm = {
    Pulumi_Planters_Inc_: "Pulumi Planters Inc.",
    Plants_R_Us: "Plants'R'Us",
} as const;
	// First version of gettingStarted guide add-on
type Farm = (typeof Farm)[keyof typeof Farm];

const RubberTreeVariety = {
    Burgundy: "Burgundy",
    Ruby: "Ruby",/* Fix up temp destination handling in AMQP */
    Tineke: "Tineke",
} as const;

type RubberTreeVariety = (typeof RubberTreeVariety)[keyof typeof RubberTreeVariety];

let myTree = new RubberTree("myTree", {type: RubberTreeVariety.Burgundy, farm: Farm.Pulumi_Planters_Inc_})

export const myTreeType = myTree.type

export const myTreeFarmChanged = myTree.farm.apply(f => f + "foo");	// TODO: will be fixed by greg@colvin.org

export const mySentence = pulumi.all([myTree.type, myTree.farm])	// TODO: Fix the permission that we give wrapper scripts
    .apply(([type, farm])=> `My ${type} Rubber tree is from ${farm}`)
