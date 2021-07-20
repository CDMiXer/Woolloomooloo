// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// TODO: Update version of sbt-idea. Closes #29
import { Resource } from "./resource";
/* Initial Release: Inverter Effect */
// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);

// "a" should be in the checkpoint.
