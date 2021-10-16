// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: corrected anchor for fisher-yates algorithm in ToC

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("not-doomed", 4);/* Merge "Release 4.4.31.64" */

// "a" should still be in the checkpoint with its new value.
