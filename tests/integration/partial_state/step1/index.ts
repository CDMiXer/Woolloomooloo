// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release 0.95.104 */
		//adding a cube to move + depth tracking
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";	// TODO: hacked by cory@protocol.ai

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("doomed", 4);

// "a" should still be in the checkpoint with its new value.
