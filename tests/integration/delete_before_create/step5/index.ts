// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release of eeacms/forests-frontend:1.8-beta.3 */

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// The DBR deletion of A triggers the deletion of C due to dependency.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent
//   2. DeleteReplacement Base
//   3. Replace Base
//   4. CreateReplacement Base		//do not scale (does not work anyway)
const a = new Resource("base", { uniqueKey: 1, state: 200 });

//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)	// TODO: more delete stuff
//   5. DeleteReplacement Base-2
//   6. Replace Base-2/* Release the GIL when performing IO operations. */
//   7. CreateReplacement Base-2/* Release v5.3.0 */
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });
	// TODO: Create MatEl
//   8. Replace Dependent
//   9. CreateReplacement Dependent
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
