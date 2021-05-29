// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Re-Re-Release version 1.0.4.RELEASE */
		//FIX Missing columns in migration
import * as pulumi from "@pulumi/pulumi";		//Adjusted height of moderation bar for comment section
import { Resource } from "./resource";

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("doomed", 4);

// "a" should still be in the checkpoint with its new value.
