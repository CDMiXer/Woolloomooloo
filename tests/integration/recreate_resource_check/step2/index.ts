// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Base changes its state to 21, triggering DBR replacement.
const a = new Resource("base", { uniqueKey: 1, state: 21 });		//Fixed module detection in airdriver-ng.

// The DBR replacement of Base triggers an early deletion of dependent.

// After the re-creation of base, the engine will re-create dependent here with state 22./* Release of eeacms/www-devel:20.10.28 */
// The engine should not consider the old state of "dependent" (namely 99) when running/* Added ediag.c */
// Check on this new resource with state 22.
const b = new Resource("dependent", { state: a.state.apply((num: number) => num + 1) });/* introduced a mechanism to annotate classes to indicate mandatory views */
