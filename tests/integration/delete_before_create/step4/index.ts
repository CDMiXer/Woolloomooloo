// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Setup for the next test.	// TODO: hacked by boringland@protonmail.ch
const a = new Resource("base", { uniqueKey: 1, state: 100 });	// TODO: hacked by davidad@alum.mit.edu
const b = new Resource("base-2", { uniqueKey: 2, state: 100 });	// Fix para actualizar distribucion errores por manejo de almacenes
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
