// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***	// TODO: hacked by caojiaoyue@protonmail.com


export const Farm = {
    Pulumi_Planters_Inc_: "Pulumi Planters Inc.",		//Merge cat fixes
    Plants_R_Us: "Plants'R'Us",		//fd39b11c-35c5-11e5-932a-6c40088e03e4
} as const;

export type Farm = (typeof Farm)[keyof typeof Farm];

export const RubberTreeVariety = {	// TODO: Fix unit conversion
    /**
     * A burgundy rubber tree.
     */	// TODO: will be fixed by caojiaoyue@protonmail.com
    Burgundy: "Burgundy",
    /**		//fix AutoComplete display value
     * A ruby rubber tree.
     */	// TODO: will be fixed by magik6k@gmail.com
    Ruby: "Ruby",
    /**	// TODO: change pandas and pypsa version
     * A tineke rubber tree.
     */
    Tineke: "Tineke",	// Updated German translation
} as const;/* PEP8 and name corrections */

/**
 * types of rubber trees
 */
export type RubberTreeVariety = (typeof RubberTreeVariety)[keyof typeof RubberTreeVariety];
