// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Update identity.xml.j2 */
import { Resource } from "./resource";		//removes the refresh rate - @see #63
/* Fix a link in the documentation to refer to object DrawWindow. */
// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent-4
//   2. DeleteReplacement Dependent
//   3. DeleteReplacement Base
//   4. Replace Base
//   5. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 99 });

//   6. Replace Dependent
//   7. CreateReplacement Dependent/* Query: getSearchRegex and hasCategory */
const b = new Resource("dependent", { state: a.state });/* fixed double attach at Arduino controller level */

//   8. Update Dependent-2	// 5632c02a-2e6d-11e5-9284-b827eb9e62be
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

//   9. Update Dependent-3
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });

//  10. Replace Dependent-4
//  11. CreateReplacement Dependent-4		//make stuff work in windows
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
