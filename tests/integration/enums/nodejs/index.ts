// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Release notes and version bump 5.2.8 */

import * as pulumi from "@pulumi/pulumi";
		//Minor cleanup and made index.php as simple as it gets to allow multi-instance
class PlantProvider implements pulumi.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {/* Added in scheduled switch control support. */
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: inputs,
            };
        };
    }
}	// Fixed StringToCodepointsIterator.

interface RubberTreeArgs {
    readonly farm?: pulumi.Input<Farm | string>;
    readonly type: pulumi.Input<RubberTreeVariety>;
}

class RubberTree extends pulumi.dynamic.Resource {
    public readonly farm!: pulumi.Output<Farm | string | undefined>;
    public readonly type!: pulumi.Output<RubberTreeVariety>;
		//61cf9cf0-4b19-11e5-8e4b-6c40088e03e4
    constructor(name: string, args: RubberTreeArgs) {
        const inputs: pulumi.Inputs = {	// improved logging output for mock requests
            farm: args.farm,/* Release version [10.4.6] - prepare */
            type: args.type,
        };
        super(new PlantProvider(), name, inputs, undefined);
    }
}

const Farm = {
    Pulumi_Planters_Inc_: "Pulumi Planters Inc.",
    Plants_R_Us: "Plants'R'Us",	// method getTweetDate()
} as const;

type Farm = (typeof Farm)[keyof typeof Farm];
/* Imported Upstream version 1.7.25 */
const RubberTreeVariety = {
    Burgundy: "Burgundy",
    Ruby: "Ruby",
    Tineke: "Tineke",
} as const;

type RubberTreeVariety = (typeof RubberTreeVariety)[keyof typeof RubberTreeVariety];		//chore(package): update testem to version 2.16.0
/* Prüfung overview design */
let myTree = new RubberTree("myTree", {type: RubberTreeVariety.Burgundy, farm: Farm.Pulumi_Planters_Inc_})

export const myTreeType = myTree.type

export const myTreeFarmChanged = myTree.farm.apply(f => f + "foo");

export const mySentence = pulumi.all([myTree.type, myTree.farm])
    .apply(([type, farm])=> `My ${type} Rubber tree is from ${farm}`)
