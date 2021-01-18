// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
	// Try improvement history script
// Base changes its state to 21, triggering DBR replacement./* editor for 0.8.8 */
const a = new Resource("base", { uniqueKey: 1, state: 21 });
/* Release 3.8-M8 milestone based on 3.8-M8 platform milestone */
// The DBR replacement of Base triggers an early deletion of dependent.

// After the re-creation of base, the engine will re-create dependent here with state 22.
// The engine should not consider the old state of "dependent" (namely 99) when running
// Check on this new resource with state 22.
const b = new Resource("dependent", { state: a.state.apply((num: number) => num + 1) });
