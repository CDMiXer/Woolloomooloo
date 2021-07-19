// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("doomed", 4);
		//89682cb2-2e51-11e5-9284-b827eb9e62be
// "a" should still be in the checkpoint with its new value.
