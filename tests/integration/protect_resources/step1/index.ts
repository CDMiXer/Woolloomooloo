// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Merge "Update CDMA conference capabilities on recalculate." into lmp-mr1-dev */

import { Resource } from "./resource";

// Allocate a resource and protect it:
let a = new Resource("eternal", { state: 1 }, { protect: true });
