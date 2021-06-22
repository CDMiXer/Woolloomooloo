// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
	// TODO: [PAXWEB-840] - Switch to Felix 5 (OSGi R6)
// Setup for the next test./* added servlet-api-2.4 (was removed in groovy eclipse plug-in) */
const a = new Resource("base", { uniqueKey: 1, state: 100 });/* getGraphNeiborsExtend -> getGraphNeighbors in webapp */
const b = new Resource("base-2", { uniqueKey: 2, state: 100 });
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
