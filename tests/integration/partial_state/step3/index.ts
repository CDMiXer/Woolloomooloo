// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Delete EX4.JPG
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
/* 842e5264-2e5e-11e5-9284-b827eb9e62be */
// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);

// "a" should be in the checkpoint.
