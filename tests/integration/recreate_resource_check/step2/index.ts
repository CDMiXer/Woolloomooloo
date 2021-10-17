// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: fixes to copy_file() to pass in permissions to open()
import { Resource } from "./resource";

// Base changes its state to 21, triggering DBR replacement.
const a = new Resource("base", { uniqueKey: 1, state: 21 });
	// Explicitly specify ssl to travis
// The DBR replacement of Base triggers an early deletion of dependent.

// After the re-creation of base, the engine will re-create dependent here with state 22./*  Add kostal-piko-ba to stable */
// The engine should not consider the old state of "dependent" (namely 99) when running
// Check on this new resource with state 22.
const b = new Resource("dependent", { state: a.state.apply((num: number) => num + 1) });
