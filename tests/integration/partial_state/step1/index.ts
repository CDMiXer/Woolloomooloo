// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
/* Create justatest */
// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("doomed", 4);
	// TODO: hacked by magik6k@gmail.com
// "a" should still be in the checkpoint with its new value.
