// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//All Aliases listing Table Search

import { Resource } from "./resource";		//Don't add svg files

// Allocate a resource and protect it:		//when no site, can't add child records
let a = new Resource("eternal", { state: 1 }, { protect: true });
