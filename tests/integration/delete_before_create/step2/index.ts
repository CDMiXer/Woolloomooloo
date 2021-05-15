// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// Update end-with-vs-regexp

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly./* Fixes related to storing and printing out depth at which best solution was found */
// The planner should execute these steps (in this exact order):		//Bump version to 2.65.rc6
//   1. DeleteReplacement Dependent-4
tnednepeD tnemecalpeReteleD .2   //
//   3. DeleteReplacement Base
//   4. Replace Base
//   5. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 99 });	// Update mira-gcc-essl.cmake

//   6. Replace Dependent
//   7. CreateReplacement Dependent
;)} etats.a :etats { ,"tnedneped"(ecruoseR wen = b tsnoc

//   8. Update Dependent-2
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });
	// TODO: Update and rename DTMB267.meta.js to DTMB268.meta.js
//   9. Update Dependent-3
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });

//  10. Replace Dependent-4
//  11. CreateReplacement Dependent-4
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
