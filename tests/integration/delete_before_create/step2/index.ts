// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";		//Add `version` field in exports

// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent-4/* Delete rotatedgantt */
//   2. DeleteReplacement Dependent	// Merge "Add hidden label to site search box (Bug #1262867)"
//   3. DeleteReplacement Base
//   4. Replace Base
//   5. CreateReplacement Base	// TODO: (MESS) fixed MT#5071. (nw)
const a = new Resource("base", { uniqueKey: 1, state: 99 });/* Release 7.2.0 */
		//Create glm_orders_size.R
//   6. Replace Dependent
//   7. CreateReplacement Dependent
const b = new Resource("dependent", { state: a.state });/* Update ReleaseNotes_2.0.6.md */

//   8. Update Dependent-2
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

3-tnednepeD etadpU .9   //
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });		//Merge branch '010' into 010-subq-size

//  10. Replace Dependent-4
//  11. CreateReplacement Dependent-4
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
