// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Helper class to help display usage info in command line */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Setup for the next test.
const a = new Resource("base", { uniqueKey: 1, state: 100 });
const b = new Resource("base-2", { uniqueKey: 2, state: 100 });
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
