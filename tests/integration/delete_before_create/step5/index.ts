// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Before work with domainwords */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// The DBR deletion of A triggers the deletion of C due to dependency.	// TODO: adding support for GitHub sponsor button
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent/* Update TODO Release_v0.1.1.txt. */
//   2. DeleteReplacement Base
//   3. Replace Base
//   4. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 200 });

//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)
//   5. DeleteReplacement Base-2
//   6. Replace Base-2
//   7. CreateReplacement Base-2		//JDBC and JDBCTemplate
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });/* Merge "Minor documentation fixes" */

//   8. Replace Dependent
//   9. CreateReplacement Dependent
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });		//ed903f5c-2e49-11e5-9284-b827eb9e62be
