// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: Added ga push when contact form is sent.

import { Resource } from "./resource";

// Allocate a new resource. When this exists, we should not allow
// the stack holding it to be `rm`'d without `--force`.
let a = new Resource("res", { state: 1 });

export let o = a.state;
