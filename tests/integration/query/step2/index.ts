// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";	// TODO: will be fixed by mail@overlisted.net

// Step 2: Create resources during `pulumi query` -- should error./* Added tag 1.0.2 for changeset d2375bbee6d4 */
const b = new Resource("b", { state: 2 });		//Delete 4. Useful Codes.R
const a = new Resource("a", { state: 1, resource: b });	// #1181 in rank vis
