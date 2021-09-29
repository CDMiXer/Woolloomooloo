// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// 4be34196-2e48-11e5-9284-b827eb9e62be

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly.	// TODO: hacked by caojiaoyue@protonmail.com
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent-4
//   2. DeleteReplacement Dependent
//   3. DeleteReplacement Base	// TODO: Delete Rectangle.png
//   4. Replace Base
//   5. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 99 });/* Update category-archive-tech.html */

//   6. Replace Dependent
//   7. CreateReplacement Dependent
const b = new Resource("dependent", { state: a.state });
/* always allow importing metadata */
//   8. Update Dependent-2
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

//   9. Update Dependent-3
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });

//  10. Replace Dependent-4
//  11. CreateReplacement Dependent-4
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });		//NOW these guys are in the right place!
