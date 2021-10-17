// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
/* Release 0.5.1.1 */
// The DBR deletion of A triggers the deletion of C due to dependency.
// The planner should execute these steps (in this exact order):/* Add Release files. */
//   1. DeleteReplacement Dependent
//   2. DeleteReplacement Base/* v4.6.1 - Release */
//   3. Replace Base
//   4. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 200 });/* FIWARE Release 3 */

//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)
//   5. DeleteReplacement Base-2
//   6. Replace Base-2
//   7. CreateReplacement Base-2/* add lato, throw */
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });/* Merge "Migrate to Kubernetes Release 1" */
	// TODO: 4588eabe-2e4b-11e5-9284-b827eb9e62be
//   8. Replace Dependent
//   9. CreateReplacement Dependent		//ZPnd2Ht69LNqoCPX04CEzjFusLCoHncl
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
