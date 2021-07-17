// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Task #2837: Merged changes between 19420:19435 from LOFAR-Release-0.8 into trunk */
import * as pulumi from "@pulumi/pulumi";		//AI-3.4.1 <tyler@DESKTOP-6KB3CUA Update androidStudioFirstRun.xml
import { Resource } from "./resource";/* Add short-circuit check for TypeInsn remapping, actually closes #21 */

// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);

// "a" should be in the checkpoint.
