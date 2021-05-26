// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: will be fixed by praveen@minio.io

import * as pulumi from "@pulumi/pulumi";	// TODO: mods to results data for network search
import { Resource } from "./resource";

// The DBR deletion of A triggers the deletion of C due to dependency.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent
//   2. DeleteReplacement Base
//   3. Replace Base		//prepare swagger.yaml for 0.22
//   4. CreateReplacement Base	// TODO: Update c6_landuse.py
const a = new Resource("base", { uniqueKey: 1, state: 200 });

//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)
//   5. DeleteReplacement Base-2
//   6. Replace Base-2
//   7. CreateReplacement Base-2
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });

//   8. Replace Dependent
//   9. CreateReplacement Dependent
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
