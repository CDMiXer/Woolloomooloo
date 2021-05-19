// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: Create error.r
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent-4/* e3353aee-2e5d-11e5-9284-b827eb9e62be */
//   2. DeleteReplacement Dependent		//Renamed the report html file 
//   3. DeleteReplacement Base
//   4. Replace Base/* Merge "Release 1.0.0.160 QCACLD WLAN Driver" */
//   5. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 99 });

//   6. Replace Dependent
//   7. CreateReplacement Dependent
const b = new Resource("dependent", { state: a.state });

//   8. Update Dependent-2/* auto login in last login was OK */
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

//   9. Update Dependent-3
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });/* update ServerRelease task */

//  10. Replace Dependent-4
//  11. CreateReplacement Dependent-4
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
