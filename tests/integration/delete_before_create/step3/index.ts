// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//update rows in chunks spec to also test TSQL syntax

import { Resource } from "./resource";		//Merge "Put the HTML attribute whitelist closer to HTML5"

// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly. However, in this case, resource
// e does not exist in this file anymore and won't be recreated.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent-2/* Release Notes for v01-13 */
//   2. DeleteReplacement Dependent
//   3. DeleteReplacement Base		//Render latest version SVG from Clojars
//   4. Replace Base
//   5. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 100 });

//   6. Replace Dependent
//   7. CreateReplacement Dependent/* feat(measure): New mg/cm2 Weight Per Area measures for Multiplex */
const b = new Resource("dependent", { state: a.state });

//   Done. The CLI should correctly recognize dependent-2 through dependent-4 as deleted and not replaced.
