// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Allocate a new resource. When this exists, we should not allow/* Changing layout, reordering components. */
// the stack holding it to be `rm`'d without `--force`.
let a = new Resource("res", { state: 1 });/* 696113a6-2e43-11e5-9284-b827eb9e62be */

export let o = a.state;
