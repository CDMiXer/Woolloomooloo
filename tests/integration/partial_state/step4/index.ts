// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";/* disambiguate [a;b]f: case [1,2,3]f of {[a;3]f -> a} works now :-) */
/* Merge "Release 3.2.3.343 Prima WLAN Driver" */
// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("not-doomed", 4);

// "a" should still be in the checkpoint with its new value.
