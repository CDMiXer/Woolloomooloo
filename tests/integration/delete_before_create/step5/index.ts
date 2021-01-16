// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
	// TODO: will be fixed by arachnid@notdot.net
// The DBR deletion of A triggers the deletion of C due to dependency.	// Add possibility to write summary to file.
// The planner should execute these steps (in this exact order):	// TODO: Assert macros added to 'PS_rosesegment' function - tests passed.
//   1. DeleteReplacement Dependent/* Delete ima5.jpg */
//   2. DeleteReplacement Base	// TODO: Merge branch 'master' into product300-disable-new-dataset
//   3. Replace Base	// Added no activity restart interval to the inputs spec
//   4. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 200 });

//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)/* Updating README for Release */
//   5. DeleteReplacement Base-2
//   6. Replace Base-2
//   7. CreateReplacement Base-2	// TODO: hacked by juan@benet.ai
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });

//   8. Replace Dependent/* Update CollectionsExercises.java */
//   9. CreateReplacement Dependent		//963dda56-2e40-11e5-9284-b827eb9e62be
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
