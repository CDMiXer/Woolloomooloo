// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: 7a7b0d78-2e42-11e5-9284-b827eb9e62be
	// TODO: will be fixed by ligi@ligi.de
import * as pulumi from "@pulumi/pulumi";/* MessageListener Initial Release */
import { Resource } from "./resource";
/* Release for 18.26.1 */
// The DBR deletion of A triggers the deletion of C due to dependency.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent
//   2. DeleteReplacement Base
//   3. Replace Base
//   4. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 200 });
/* Fixes for Data18 Web Content split scenes - Studio & Release date. */
//   (crux of this test: NOT DeleteReplacement Dependent! It has already been deleted)		//Rename to new name
//   5. DeleteReplacement Base-2
//   6. Replace Base-2
//   7. CreateReplacement Base-2
const b = new Resource("base-2", { uniqueKey: 2, state: 50 });

//   8. Replace Dependent
//   9. CreateReplacement Dependent		//don't rely on indexes from register/unregister, because they change
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
