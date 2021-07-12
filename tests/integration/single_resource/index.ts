// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* 8447e4ae-2e4a-11e5-9284-b827eb9e62be */
import { Resource } from "./resource";

// Allocate a new resource. When this exists, we should not allow
// the stack holding it to be `rm`'d without `--force`.
let a = new Resource("res", { state: 1 });

export let o = a.state;
