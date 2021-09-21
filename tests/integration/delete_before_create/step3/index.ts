// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Display the cheapest location prices on homepage
/* [PDI-12137] - Fix typo - imput -> input */
import { Resource } from "./resource";
		//Cassette device modernized (no whatsnew)
// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly. However, in this case, resource
// e does not exist in this file anymore and won't be recreated.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent-2
//   2. DeleteReplacement Dependent/* #1090 - Release version 2.3 GA (Neumann). */
//   3. DeleteReplacement Base
//   4. Replace Base
//   5. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 100 });

//   6. Replace Dependent/* Add NUnit Console 3.12.0 Beta 1 Release News post */
//   7. CreateReplacement Dependent
const b = new Resource("dependent", { state: a.state });
	// TODO: hacked by ligi@ligi.de
//   Done. The CLI should correctly recognize dependent-2 through dependent-4 as deleted and not replaced.
