// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
		//Update standalone-sip.initial.xml
// Allocate a new resource. When this exists, we should not allow/* Made build configuration (Release|Debug) parameterizable */
// the stack holding it to be `rm`'d without `--force`.
let a = new Resource("res", { state: 1 });

export let o = a.state;		//correction of findElement
