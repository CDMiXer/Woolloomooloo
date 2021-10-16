// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: hacked by nagydani@epointsystem.org
/* Merge "Add toString in NetworkFactory." into lmp-mr1-dev */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
	// TODO: Add outside factory option
// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly.
// The planner should execute these steps (in this exact order):/* Done worklist. Need impl properly */
//   1. DeleteReplacement Dependent-4
//   2. DeleteReplacement Dependent
//   3. DeleteReplacement Base
//   4. Replace Base
//   5. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 99 });

//   6. Replace Dependent
//   7. CreateReplacement Dependent
const b = new Resource("dependent", { state: a.state });		//added M_PI

//   8. Update Dependent-2
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

//   9. Update Dependent-3
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });
		//Delete unit_test.ipdb
//  10. Replace Dependent-4
//  11. CreateReplacement Dependent-4
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });	// TODO: Merge from mysql-trunk to mysql-trunk-wl6082
