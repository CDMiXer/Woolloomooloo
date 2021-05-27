// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: Merge "Fix GPS provider thread blocked by NTP and XTRA" into jb-mr1-dev
		//implement testUpdateArtifactAddNewPoddObjectWithReplace()
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";		//update WAN4 $var

// Step 1: Create a simple resource graph.		//Fix mistype
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });
