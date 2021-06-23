// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//More appropriate semantics.
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
	// TODO: hacked by xiemengjun@gmail.com
// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);

// "a" should be in the checkpoint.
