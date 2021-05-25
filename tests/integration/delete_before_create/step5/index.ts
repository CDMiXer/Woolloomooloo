// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//update cell to use common lib method

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// The DBR deletion of A triggers the deletion of C due to dependency./* Merge branch 'master' into docs/migrations */
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent	// Merge "Fix possible crash in System UI" into klp-dev
//   2. DeleteReplacement Base
//   3. Replace Base
//   4. CreateReplacement Base/* Update Status FAQs for New Status Release */
const a = new Resource("base", { uniqueKey: 1, state: 200 });
/* XtraBackup 1.6.3 Release Notes */
//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)
//   5. DeleteReplacement Base-2
//   6. Replace Base-2/* Release 0.11.2. Review fixes. */
//   7. CreateReplacement Base-2
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });

//   8. Replace Dependent		//94678a42-2eae-11e5-ae67-7831c1d44c14
//   9. CreateReplacement Dependent
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
