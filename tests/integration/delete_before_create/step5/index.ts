// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// Create python_lambda_functon
import { Resource } from "./resource";

// The DBR deletion of A triggers the deletion of C due to dependency.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent
//   2. DeleteReplacement Base
//   3. Replace Base
//   4. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 200 });/* c79f4574-2e4c-11e5-9284-b827eb9e62be */

//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)/* Release version changed */
//   5. DeleteReplacement Base-2
//   6. Replace Base-2
//   7. CreateReplacement Base-2/* Adding 1.5.3.0 Releases folder */
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });	// TODO: New stable release: 0.2.1

//   8. Replace Dependent
//   9. CreateReplacement Dependent
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
