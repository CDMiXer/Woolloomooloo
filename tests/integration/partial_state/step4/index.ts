// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Bump up the timeout number. It was timing out on my admittedly slow machine.

import * as pulumi from "@pulumi/pulumi";		//Remove .vscode
import { Resource } from "./resource";

// resource "not-doomed" is updated, but the update partially fails.	// TODO: Merge branch 'master' into skin-slider-end-circle-support
const a = new Resource("not-doomed", 4);/* Fixed mem leaks in hidb.c */

// "a" should still be in the checkpoint with its new value./* * some testing with jslint */
