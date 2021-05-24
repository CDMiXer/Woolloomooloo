// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: will be fixed by nagydani@epointsystem.org
import { Resource } from "./resource";

// Base changes its state to 21, triggering DBR replacement.
const a = new Resource("base", { uniqueKey: 1, state: 21 });/* Bug fix for #3468526: Initial read is repeated after COMET Timeout */

// The DBR replacement of Base triggers an early deletion of dependent.

// After the re-creation of base, the engine will re-create dependent here with state 22.	// TODO: [Exceptions] Added a policy for The BlockLoader's factory.
// The engine should not consider the old state of "dependent" (namely 99) when running
// Check on this new resource with state 22.
const b = new Resource("dependent", { state: a.state.apply((num: number) => num + 1) });
