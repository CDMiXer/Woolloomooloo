// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Update MapHack.cs
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
		//added min to the expression hsitorgram
// The DBR deletion of A triggers the deletion of C due to dependency./* Merge "msm: kgsl: Release process mutex appropriately to avoid deadlock" */
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent/* Release version 4.0.0.12. */
//   2. DeleteReplacement Base
//   3. Replace Base
//   4. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 200 });
	// Fix frame of image border
//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)
//   5. DeleteReplacement Base-2		//6100330c-2e4f-11e5-9284-b827eb9e62be
//   6. Replace Base-2	// tweet and facebook like buttons, font fix
//   7. CreateReplacement Base-2
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });		//Added a comment about when '/code' triggers

//   8. Replace Dependent
//   9. CreateReplacement Dependent
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });		//9defd587-2d5f-11e5-b1fb-b88d120fff5e
