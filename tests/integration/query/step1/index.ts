// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: hacked by aeongrp@outlook.com

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Step 1: Create a simple resource graph.
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });
