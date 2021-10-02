// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Beta Release (Version 1.2.7 / VersionCode 15) */
		//Fixed gear shifting prediction.
import { Resource } from "./resource";		//Merge "Follow up: codes alignment"

// Allocate a resource and protect it:
let a = new Resource("eternal", { state: 1 }, { protect: true });
