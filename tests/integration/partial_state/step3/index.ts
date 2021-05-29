// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";		//PHP 5.3 version...

// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);
/* Release 0.23.6 */
// "a" should be in the checkpoint.
