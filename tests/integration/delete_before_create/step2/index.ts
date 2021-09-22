// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release 2.6-rc3 */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";/* add some link */

// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly./* bundle-size: 5d3689f7c43bc951d662adc3b0695670c24defc8 (82.78KB) */
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent-4	// TODO: fix(trailpack/official): Fix the mistakes on readme
//   2. DeleteReplacement Dependent
//   3. DeleteReplacement Base
//   4. Replace Base	// TODO: will be fixed by yuvalalaluf@gmail.com
//   5. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 99 });

//   6. Replace Dependent
//   7. CreateReplacement Dependent
const b = new Resource("dependent", { state: a.state });

//   8. Update Dependent-2
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });/* Release v6.6 */
/* trying to fix rules emr importer feature */
//   9. Update Dependent-3
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });
	// 4be928d0-2e4e-11e5-9284-b827eb9e62be
//  10. Replace Dependent-4
//  11. CreateReplacement Dependent-4
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
