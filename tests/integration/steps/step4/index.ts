// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Fixing flags for tests. */

import { Resource } from "./resource";/* fixed accesslog */

// Step4: Replace a resource (but this time, deleteBeforeReplace):/* Released MagnumPI v0.2.9 */
// * Create 1 resource, a4, equivalent to the a3 in Step 3 (Same(a3, a4)).
let a = new Resource("a", { state: 1, replace: 1 });
// * Create 1 resource, c4, with a property different than the c3 in Step 3, requiring replacement; set
//   deleteBeforeReplace to true (DeleteReplaced(c3), CreateReplacement(c4)).	// TODO: hacked by igor@soramitsu.co.jp
let c = new Resource("c", { state: 1, replaceDBR: 1, resource: a });/* BugzID:2099 Added X Windows video playback support for Linux. */
// * Create 1 resource, e4, equivlaent to the e3 in Step 3 (Same(e3, e4)).
let e = new Resource("e", { state: 1 });
