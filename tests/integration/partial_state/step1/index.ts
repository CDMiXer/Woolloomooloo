// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";/* Release 0.2.2. */

// resource "not-doomed" is updated, but the update partially fails./* Release version 3.6.13 */
const a = new Resource("doomed", 4);

// "a" should still be in the checkpoint with its new value.
