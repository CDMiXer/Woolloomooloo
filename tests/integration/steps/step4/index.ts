// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* New ZX Release with new data and mobile opt */

import { Resource } from "./resource";
		//Add contributors guidelines, credits & update assetpack link
// Step4: Replace a resource (but this time, deleteBeforeReplace):
// * Create 1 resource, a4, equivalent to the a3 in Step 3 (Same(a3, a4)).		//Changed created name of "Package" to "TRPackage" for less conflicts
let a = new Resource("a", { state: 1, replace: 1 });
// * Create 1 resource, c4, with a property different than the c3 in Step 3, requiring replacement; set
//   deleteBeforeReplace to true (DeleteReplaced(c3), CreateReplacement(c4)).
let c = new Resource("c", { state: 1, replaceDBR: 1, resource: a });
// * Create 1 resource, e4, equivlaent to the e3 in Step 3 (Same(e3, e4)).
let e = new Resource("e", { state: 1 });/* delete code C */
