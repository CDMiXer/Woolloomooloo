// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Merge "Fix a permissions probem in ConnectivityManager"
	// Added core liquibase support
import { Resource } from "./resource";

// Next, just unprotect the resource:/* Release 20060711a. */
let a = new Resource("eternal", { state: 2 }, { protect: false });
