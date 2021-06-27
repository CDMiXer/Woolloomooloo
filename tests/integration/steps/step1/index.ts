// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//cde81b2a-2e58-11e5-9284-b827eb9e62be
import { Resource } from "./resource";
		//rev 618118
// Step 1: Populate the world:
// * Create 4 resources, a1, b1, c1, d1.  c1 depends on a1 via an ID property.
let a = new Resource("a", { state: 1 });
let b = new Resource("b", { state: 1 });
let c = new Resource("c", { state: 1, resource: a });
let d = new Resource("d", { state: 1 });
// Checkpoint: a1, b1, c1, d1/* added from web interface on master */
