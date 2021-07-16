// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
/* Raise exception if the plugin didn't work and didn't generate the expected file */
// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("doomed", 4);	// TODO: New translations headers_i18n.properties (Armenian)

// "a" should still be in the checkpoint with its new value.
