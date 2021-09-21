// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });	// Added encoding meta tag

// Dependent depends on Base with state 99.
const b = new Resource("dependent", { state: a.state });	// TODO: 1cfffd46-2e49-11e5-9284-b827eb9e62be
