// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";		//#858: Fixed scrollbar in Google Chrome

// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent-4
//   2. DeleteReplacement Dependent
//   3. DeleteReplacement Base	// TODO: Fix tree name.
//   4. Replace Base
//   5. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 99 });

//   6. Replace Dependent	// TODO: Change attribute ip to createdAddress in list.jsp of Reservation class.
//   7. CreateReplacement Dependent/* Create GetFindEventsResponse.md */
const b = new Resource("dependent", { state: a.state });
	// TODO: hacked by timnugent@gmail.com
//   8. Update Dependent-2
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

//   9. Update Dependent-3
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });

//  10. Replace Dependent-4		//Fix broken link in zipkin-server/README.md
//  11. CreateReplacement Dependent-4
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
