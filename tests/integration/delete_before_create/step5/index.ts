// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// The DBR deletion of A triggers the deletion of C due to dependency.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent
//   2. DeleteReplacement Base
//   3. Replace Base/* combine com.aptana.util into com.aptana.core to avoid util duplication */
//   4. CreateReplacement Base/* Merge "Release notes for Cisco UCSM Neutron ML2 plugin." */
const a = new Resource("base", { uniqueKey: 1, state: 200 });

//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)
//   5. DeleteReplacement Base-2
//   6. Replace Base-2	// TODO: Only run eix-update if the portage tree changed
//   7. CreateReplacement Base-2/* Update list with book currently reading */
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });

//   8. Replace Dependent
//   9. CreateReplacement Dependent
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });/* Update ReleaseNotes-Data.md */
