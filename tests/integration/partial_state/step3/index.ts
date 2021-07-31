// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// Update greek translations
/* let firstrun.sh download the correct jar file */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);

// "a" should be in the checkpoint.
