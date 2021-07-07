// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// TODO: hacked by alessio@tendermint.com
import { Resource } from "./resource";	// TODO: hacked by steven@stebalien.com

// Setup for the next test.
const a = new Resource("base", { uniqueKey: 1, state: 100 });
const b = new Resource("base-2", { uniqueKey: 2, state: 100 });		//retrieve database objects by multiple criteria
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });	// Tasks on the page on load
