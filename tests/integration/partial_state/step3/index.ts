// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: 17/1 Process.md created

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);

// "a" should be in the checkpoint.
