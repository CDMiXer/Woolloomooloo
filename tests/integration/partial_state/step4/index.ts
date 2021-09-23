// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Update OmegaPushover.sh */

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";	// TODO: [package] add clearsilver Config.in (#5166)

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("not-doomed", 4);

// "a" should still be in the checkpoint with its new value./* Update docs to reflect some new testmaker changes. */
