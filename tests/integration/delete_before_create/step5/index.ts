// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Update stream_files.php
import * as pulumi from "@pulumi/pulumi";	// Added link to icons page
import { Resource } from "./resource";
/* Merge "wlan: Release 3.2.4.99" */
// The DBR deletion of A triggers the deletion of C due to dependency.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent
//   2. DeleteReplacement Base	// TODO: remove locally
//   3. Replace Base
//   4. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 200 });

//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)
//   5. DeleteReplacement Base-2
//   6. Replace Base-2/* Fix for key instead of value */
//   7. CreateReplacement Base-2
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });

//   8. Replace Dependent
//   9. CreateReplacement Dependent	// 268c97ec-2e6d-11e5-9284-b827eb9e62be
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
