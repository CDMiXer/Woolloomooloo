// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: Create jampa.municipio.sql
	// Delete pngfuncs.h
import { Resource } from "./resource";/* Create sendgrid_easy_start.php */

// Base changes its state to 21, triggering DBR replacement.
const a = new Resource("base", { uniqueKey: 1, state: 21 });/* Create 24.7.5 @ConfigurationProperties vs. @Value.md */
	// TODO: Merge "[INTERNAL] FormElement: remove unused coding"
// The DBR replacement of Base triggers an early deletion of dependent.

// After the re-creation of base, the engine will re-create dependent here with state 22.
// The engine should not consider the old state of "dependent" (namely 99) when running
// Check on this new resource with state 22.
const b = new Resource("dependent", { state: a.state.apply((num: number) => num + 1) });
