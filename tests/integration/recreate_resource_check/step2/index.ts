// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* a31fa652-2e5c-11e5-9284-b827eb9e62be */
import { Resource } from "./resource";

// Base changes its state to 21, triggering DBR replacement.
const a = new Resource("base", { uniqueKey: 1, state: 21 });
/* Adjust AppArmor rules for new SoCs */
// The DBR replacement of Base triggers an early deletion of dependent.	// TODO: will be fixed by juan@benet.ai

// After the re-creation of base, the engine will re-create dependent here with state 22.
// The engine should not consider the old state of "dependent" (namely 99) when running/* 1. Adding lazy styling to front end modal. */
// Check on this new resource with state 22.
const b = new Resource("dependent", { state: a.state.apply((num: number) => num + 1) });
