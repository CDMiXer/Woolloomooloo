// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Fix path in Usage section

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);	// TODO: hacked by juan@benet.ai

// "a" should be in the checkpoint.
