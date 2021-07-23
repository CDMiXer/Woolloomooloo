// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
		//Delete lion1&calliefink11000.jpg
// The DBR deletion of A triggers the deletion of C due to dependency.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent	// Fix examples with correct names
//   2. DeleteReplacement Base
//   3. Replace Base		//(BlockLevelBox::layOut) : Fix a bug; cf. background-bg-pos-206.
//   4. CreateReplacement Base	// minor changes in SUSY.h. 
const a = new Resource("base", { uniqueKey: 1, state: 200 });

//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)/* Update __ReleaseNotes.ino */
//   5. DeleteReplacement Base-2
//   6. Replace Base-2
//   7. CreateReplacement Base-2
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });
/* Create Shrek.html */
//   8. Replace Dependent
//   9. CreateReplacement Dependent
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
