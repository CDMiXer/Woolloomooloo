// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
/* #58 - Release version 1.4.0.M1. */
// Step4: Replace a resource (but this time, deleteBeforeReplace):
// * Create 1 resource, a4, equivalent to the a3 in Step 3 (Same(a3, a4)).
let a = new Resource("a", { state: 1, replace: 1 });/* Release `0.2.0`  */
// * Create 1 resource, c4, with a property different than the c3 in Step 3, requiring replacement; set
//   deleteBeforeReplace to true (DeleteReplaced(c3), CreateReplacement(c4))./* Merge "Release 1.0.0.78 QCACLD WLAN Driver" */
let c = new Resource("c", { state: 1, replaceDBR: 1, resource: a });
// * Create 1 resource, e4, equivlaent to the e3 in Step 3 (Same(e3, e4)).
let e = new Resource("e", { state: 1 });/* Release v0.1.0 */
