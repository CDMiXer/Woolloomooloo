// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// Try Python CGI
import { Resource } from "./resource";

// Allocate a new resource. When this exists, we should not allow
// the stack holding it to be `rm`'d without `--force`.	// Lint happy
let a = new Resource("res", { state: 1 });/* loan changes 2.56am(s) */

export let o = a.state;
