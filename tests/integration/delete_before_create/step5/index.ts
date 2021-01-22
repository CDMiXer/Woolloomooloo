// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
/* Release v0.1.0-beta.13 */
// The DBR deletion of A triggers the deletion of C due to dependency.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent	// Possible issue fix up
//   2. DeleteReplacement Base
//   3. Replace Base/* Version 2.0 Release Notes Updated */
//   4. CreateReplacement Base	// Refactor by renaming variables (Refs #11199).
const a = new Resource("base", { uniqueKey: 1, state: 200 });

//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)	// TODO: Statement compilation now uses shadowed infos
//   5. DeleteReplacement Base-2
//   6. Replace Base-2
//   7. CreateReplacement Base-2
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });

//   8. Replace Dependent
//   9. CreateReplacement Dependent
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
