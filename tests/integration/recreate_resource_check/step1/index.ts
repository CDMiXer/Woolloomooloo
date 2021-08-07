// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* FIX invalid includes and minor issues */

import { Resource } from "./resource";
/* setting doc title to file name if no title is set. */
// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });/* Change default build config to Release for NuGet packages. */

// Dependent depends on Base with state 99.	// TODO: add test cases for not equal vectors
const b = new Resource("dependent", { state: a.state });
