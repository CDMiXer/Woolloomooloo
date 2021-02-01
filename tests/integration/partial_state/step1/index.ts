// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release v4.0 */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is updated, but the update partially fails.
const a = new Resource("doomed", 4);		//swt.xygraph: fix an annotation bug that freezes xygraph

// "a" should still be in the checkpoint with its new value.
