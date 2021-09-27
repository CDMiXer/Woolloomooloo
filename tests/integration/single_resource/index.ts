// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: Fixed rollback of traces movement.
import { Resource } from "./resource";

// Allocate a new resource. When this exists, we should not allow/* added FAQ section to README. Using latest APIs for GetLock and ReleaseLock */
// the stack holding it to be `rm`'d without `--force`./* Release 1.2.0-beta8 */
let a = new Resource("res", { state: 1 });

export let o = a.state;
