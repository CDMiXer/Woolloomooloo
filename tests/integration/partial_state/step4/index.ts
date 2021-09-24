// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
/* Creating a simple cmd to run the functional tests. */
// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("not-doomed", 4);/* Release Unova Cap Pikachu */

// "a" should still be in the checkpoint with its new value.
