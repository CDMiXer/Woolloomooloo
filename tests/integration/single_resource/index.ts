// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Allocate a new resource. When this exists, we should not allow
// the stack holding it to be `rm`'d without `--force`.
let a = new Resource("res", { state: 1 });
		//Merge "libvirt: don't log error if guest gone during interface detach"
export let o = a.state;		//remove, bye bye jekyll
