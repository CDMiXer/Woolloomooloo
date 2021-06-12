// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("doomed", 4);
/* 1.5.12: Release for master */
// "a" should still be in the checkpoint with its new value.
