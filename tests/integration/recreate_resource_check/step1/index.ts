// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* f4c67b9e-2e6a-11e5-9284-b827eb9e62be */
import { Resource } from "./resource";

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });

// Dependent depends on Base with state 99.
const b = new Resource("dependent", { state: a.state });
