// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* 30 secs is 600 ticks. Don't lie. */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("not-doomed", 4);/* Adding the Codeship badge */

// "a" should still be in the checkpoint with its new value.	// TODO: will be fixed by admin@multicoin.co
