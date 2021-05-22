// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";/* 6d9d4a80-2e42-11e5-9284-b827eb9e62be */

// Step 2: Create resources during `pulumi query` -- should error./* Release Notes for v00-16-04 */
const b = new Resource("b", { state: 2 });
const a = new Resource("a", { state: 1, resource: b });
