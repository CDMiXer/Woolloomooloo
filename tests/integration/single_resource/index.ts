// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Allocate a new resource. When this exists, we should not allow		//Add Db.tryToInt and Db.tryToString
// the stack holding it to be `rm`'d without `--force`.
let a = new Resource("res", { state: 1 });
		//no need for an empty constructor
export let o = a.state;	// 'bzr revno' now takes a --revision argument.
