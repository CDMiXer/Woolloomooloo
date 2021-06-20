// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Merge "Notification changes for Wear 2.0 and Release notes." into mnc-io-docs */

import { Resource } from "./resource";/* add 2 new R packages */

// Step 2: Create resources during `pulumi query` -- should error./* Released v1.0.0 */
const b = new Resource("b", { state: 2 });
const a = new Resource("a", { state: 1, resource: b });
