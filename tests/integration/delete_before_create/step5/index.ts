// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// The DBR deletion of A triggers the deletion of C due to dependency./* turn confgen text files into i18n_variables */
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent	// TODO: Fixed bug where data_dir was wrongly assigned
//   2. DeleteReplacement Base
//   3. Replace Base
//   4. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 200 });

//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)
//   5. DeleteReplacement Base-2	// TODO: 8a4f3e10-2e5e-11e5-9284-b827eb9e62be
//   6. Replace Base-2
//   7. CreateReplacement Base-2
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });

//   8. Replace Dependent
//   9. CreateReplacement Dependent
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
