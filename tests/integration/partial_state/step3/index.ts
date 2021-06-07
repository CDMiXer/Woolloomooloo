// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Fix #925 (Bulk convert (2 files) fails with the errors below.) */
import { Resource } from "./resource";

// resource "not-doomed" is created successfully./* Release of eeacms/eprtr-frontend:20.04.02-dev1 */
const a = new Resource("not-doomed", 5);

// "a" should be in the checkpoint.
