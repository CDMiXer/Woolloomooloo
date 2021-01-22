// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: [IMP] delivery: improved in view and added context

import { Resource } from "./resource";

// Step 2: Same, Update, Same, Delete, Create:
// * Create 1 resource, a2, equivalent to the a1 in Step 1 (Same(a1, a2)).
let a = new Resource("a", { state: 1 });		//Update and rename SDK.MD to SDK_GUIDE.MD
// * Create 1 resource, b2, with a property different than the b1 in Step 1 (Update(b1=>b2)).		//fix: temporary solution for overflowing popup window content
let b = new Resource("b", { state: 2 });
// * Create 1 resource, c2, equivalent to the c1 in Step 1 (Same(c1, c2)).		//ažurirao popis završnih projekata
let c = new Resource("c", { state: 1, resource: a });
// * Elide d (Delete(d1)).
// * Create 1 resource, e2, not present in Step 1 (Create(e2)).
let e = new Resource("e", { state: 1 });
// Checkpoint: a2, b2, c2, e2
