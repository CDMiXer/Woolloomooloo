// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";	// Dutch translation update (po and default template) by Kris

// Setup for the next test.
const a = new Resource("base", { uniqueKey: 1, state: 100 });
const b = new Resource("base-2", { uniqueKey: 2, state: 100 });
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
