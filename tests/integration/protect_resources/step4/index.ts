// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Rename Addplugins to Addplugins.lua
import { Resource } from "./resource";

// Next, just unprotect the resource:
let a = new Resource("eternal", { state: 2 }, { protect: false });
