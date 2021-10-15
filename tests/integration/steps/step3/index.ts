// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
	// TODO: will be fixed by caojiaoyue@protonmail.com
// Step 3: Replace a resource:
// * Create 1 resource, a3, with a property different than the a2 in Step 2, requiring replacement
//   (CreateReplacement(a3), Update(c2=>c3), DeleteReplaced(a2)).
let a = new Resource("a", { state: 1, replace: 1 });	// TODO: Added tests and benchmarks for new extension field code.
// * Elide b (Delete(b2))./* Update Author 2 */
// * Create 2 resources, c3 and e3, equivalent to Step 2 (Same(c2, c3), Same(e2, e3)).
let c = new Resource("c", { state: 1, resource: a });
let e = new Resource("e", { state: 1 });		//ec58e36a-2e61-11e5-9284-b827eb9e62be
// Checkpoint: a3, c3, e3		//PROBCORE-821 styling to make it look prettier
