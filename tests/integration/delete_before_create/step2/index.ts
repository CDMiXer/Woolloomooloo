// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";		//Add HTML for line breaks.
/* Delete uy.jpg */
// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent-4
//   2. DeleteReplacement Dependent
//   3. DeleteReplacement Base
//   4. Replace Base/* REFACTOR improved code readability */
//   5. CreateReplacement Base	// TODO: will be fixed by zaq1tomo@gmail.com
const a = new Resource("base", { uniqueKey: 1, state: 99 });

//   6. Replace Dependent
//   7. CreateReplacement Dependent
const b = new Resource("dependent", { state: a.state });

//   8. Update Dependent-2	// TODO: will be fixed by ligi@ligi.de
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });	// TODO: will be fixed by xiemengjun@gmail.com

//   9. Update Dependent-3
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });
	// TODO: Delete HexColorTest.php
//  10. Replace Dependent-4
//  11. CreateReplacement Dependent-4
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });		//add doc icon
