// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: will be fixed by fjl@ethereum.org
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("not-doomed", 4);

// "a" should still be in the checkpoint with its new value.
