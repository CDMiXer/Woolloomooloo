// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
		//Merge branch 'master' into TIMOB-20024
// Allocate a new resource. When this exists, we should not allow
// the stack holding it to be `rm`'d without `--force`.
let a = new Resource("res", { state: 1 });
	// TODO: will be fixed by magik6k@gmail.com
export let o = a.state;
