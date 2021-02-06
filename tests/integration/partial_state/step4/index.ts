// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* 75668e54-2e4a-11e5-9284-b827eb9e62be */

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("not-doomed", 4);

// "a" should still be in the checkpoint with its new value.		//Update and rename Index to Contents.md
