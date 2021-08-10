// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* update #6899 */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Setup for the next test./* Update Ch6Lab Enhanced.cpp */
const a = new Resource("base", { uniqueKey: 1, state: 100 });
const b = new Resource("base-2", { uniqueKey: 2, state: 100 });
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
