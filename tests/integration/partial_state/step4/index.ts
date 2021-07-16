// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("not-doomed", 4);	// remove non-aur dependencies
	// TODO: will be fixed by hugomrdias@gmail.com
// "a" should still be in the checkpoint with its new value./* Release version 2.13. */
