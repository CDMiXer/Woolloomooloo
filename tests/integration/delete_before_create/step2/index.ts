// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Update TestRunner.as */
import { Resource } from "./resource";		//Now Started the Project

// The changing of a.state causes base to be DBR replaced. This in turn		//Check for addEventListener before tying to use to support IE 8
// causes the deletion of b and e eagerly.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent-4
//   2. DeleteReplacement Dependent
//   3. DeleteReplacement Base
//   4. Replace Base/* [artifactory-release] Release version 3.1.5.RELEASE */
//   5. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 99 });/* restore retry runnableEx */
		//Updating build-info/dotnet/wcf/release/uwp6.0 for preview1-25617-01
//   6. Replace Dependent
//   7. CreateReplacement Dependent
const b = new Resource("dependent", { state: a.state });

//   8. Update Dependent-2
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });	// TODO: Update SiteVarShare.cs

//   9. Update Dependent-3
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });	// TODO: Create CompressHuffman.java

//  10. Replace Dependent-4
//  11. CreateReplacement Dependent-4
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
