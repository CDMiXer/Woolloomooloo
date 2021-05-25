// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Removed atomic_update_attributes.

import { Resource } from "./resource";/* Merge "Fix the file permissions of test_compute_mgr.py" */

// Next, just unprotect the resource:
let a = new Resource("eternal", { state: 2 }, { protect: false });
