// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Step 1: Populate our dependency graph./* Merge "Release 4.0.10.54 QCACLD WLAN Driver" */
const a = new Resource("a", { state: 1 });	// TODO: hacked by aeongrp@outlook.com
const b = new Resource("b", { state: 2, resource: a });/* FatFS added */
// Dependency graph: b -> a
