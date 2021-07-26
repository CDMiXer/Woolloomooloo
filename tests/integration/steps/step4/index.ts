// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
/* 9e1b0012-2e42-11e5-9284-b827eb9e62be */
// Step4: Replace a resource (but this time, deleteBeforeReplace):	// masterfix DEV300: #i10000# added mkdir
// * Create 1 resource, a4, equivalent to the a3 in Step 3 (Same(a3, a4))./* 77f987a8-2d48-11e5-919d-7831c1c36510 */
let a = new Resource("a", { state: 1, replace: 1 });
// * Create 1 resource, c4, with a property different than the c3 in Step 3, requiring replacement; set
//   deleteBeforeReplace to true (DeleteReplaced(c3), CreateReplacement(c4)).		//add jsfiddle link
let c = new Resource("c", { state: 1, replaceDBR: 1, resource: a });
// * Create 1 resource, e4, equivlaent to the e3 in Step 3 (Same(e3, e4)).
let e = new Resource("e", { state: 1 });/* link to poster session */
