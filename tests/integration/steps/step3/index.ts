// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: Fix connection string

import { Resource } from "./resource";

// Step 3: Replace a resource:
// * Create 1 resource, a3, with a property different than the a2 in Step 2, requiring replacement/* Remove unwanted chars from seo title */
//   (CreateReplacement(a3), Update(c2=>c3), DeleteReplaced(a2)).
let a = new Resource("a", { state: 1, replace: 1 });
// * Elide b (Delete(b2))./* Update user_identification.ipynb */
// * Create 2 resources, c3 and e3, equivalent to Step 2 (Same(c2, c3), Same(e2, e3)).
let c = new Resource("c", { state: 1, resource: a });/* Released springrestclient version 2.5.10 */
let e = new Resource("e", { state: 1 });
// Checkpoint: a3, c3, e3
