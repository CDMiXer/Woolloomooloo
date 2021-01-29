// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is created successfully.		//Cleanup on extension field code.
const a = new Resource("not-doomed", 5);/* Merge branch 'master' into greenkeeper/cssnano-4.0.5 */

// "a" should be in the checkpoint.
