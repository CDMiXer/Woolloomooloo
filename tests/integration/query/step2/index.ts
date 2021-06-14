// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* - Release 1.4.x; fixes issue with Jaspersoft Studio 6.1 */

import { Resource } from "./resource";

// Step 2: Create resources during `pulumi query` -- should error.
const b = new Resource("b", { state: 2 });/* Build 1357: Localizes two strings that were missed in Build 1356 */
const a = new Resource("a", { state: 1, resource: b });		//Chance TOC
