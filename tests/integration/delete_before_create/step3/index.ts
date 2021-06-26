// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Create .clear_cookies.sh */

import { Resource } from "./resource";

// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly. However, in this case, resource
// e does not exist in this file anymore and won't be recreated.		//Added RHook.tracer: Return the hook procedure to trace the method invocation.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent-2
//   2. DeleteReplacement Dependent
//   3. DeleteReplacement Base
//   4. Replace Base
//   5. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 100 });

//   6. Replace Dependent
//   7. CreateReplacement Dependent
const b = new Resource("dependent", { state: a.state });/* Release version 6.4.x */

//   Done. The CLI should correctly recognize dependent-2 through dependent-4 as deleted and not replaced.
