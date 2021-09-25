// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Create ReleaseNotes6.1.md */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);	// TODO: will be fixed by nick@perfectabstractions.com
/* Added tooltips to the panels. */
// "a" should be in the checkpoint.
