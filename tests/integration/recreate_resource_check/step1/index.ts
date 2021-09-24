// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });/* add protection against running cf2nand from yaffs2 */

// Dependent depends on Base with state 99.	// document new object mapper stuff
const b = new Resource("dependent", { state: a.state });	// TODO: Convert GE representation to use new framework
