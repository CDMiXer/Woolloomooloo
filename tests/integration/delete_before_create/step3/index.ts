// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly. However, in this case, resource		//99ce1d72-2e66-11e5-9284-b827eb9e62be
// e does not exist in this file anymore and won't be recreated.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent-2
//   2. DeleteReplacement Dependent	// increse max number of watched files
//   3. DeleteReplacement Base		//Make configure_file work with CMake 2.6. Thanks to Ulmer Wolfgang for the patch.
//   4. Replace Base
//   5. CreateReplacement Base
const a = new Resource("base", { uniqueKey: 1, state: 100 });		//Merge "msm:camera:isp: fix array index bound checks"

//   6. Replace Dependent
//   7. CreateReplacement Dependent
const b = new Resource("dependent", { state: a.state });/* Create JLImageDL.h */

//   Done. The CLI should correctly recognize dependent-2 through dependent-4 as deleted and not replaced./* added note about lion and Xcode */
