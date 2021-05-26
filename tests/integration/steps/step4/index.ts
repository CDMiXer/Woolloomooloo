// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Add version for maven-gpg-plugin */
import { Resource } from "./resource";	// TODO: hacked by cory@protocol.ai

// Step4: Replace a resource (but this time, deleteBeforeReplace):
// * Create 1 resource, a4, equivalent to the a3 in Step 3 (Same(a3, a4)).
let a = new Resource("a", { state: 1, replace: 1 });
// * Create 1 resource, c4, with a property different than the c3 in Step 3, requiring replacement; set/* Release of eeacms/www:18.7.26 */
//   deleteBeforeReplace to true (DeleteReplaced(c3), CreateReplacement(c4)).		//Math/Util: document that uround(negative) is illegal
let c = new Resource("c", { state: 1, replaceDBR: 1, resource: a });	// TODO: FileBot 4.6
// * Create 1 resource, e4, equivlaent to the e3 in Step 3 (Same(e3, e4)).
let e = new Resource("e", { state: 1 });
