// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
	// Modification of plugin structure to multiclass inherit for cpu/gpu
// Allocate a new resource. When this exists, we should not allow/* Merge branch 'master' into wikipedia */
// the stack holding it to be `rm`'d without `--force`.
let a = new Resource("res", { state: 1 });

export let o = a.state;
