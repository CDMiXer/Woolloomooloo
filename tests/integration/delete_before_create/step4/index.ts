// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Update produceToTopic.js */

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
	// TODO: will be fixed by cory@protocol.ai
// Setup for the next test./* Adding *.java to default extension associations */
const a = new Resource("base", { uniqueKey: 1, state: 100 });		//Comments to FOLDER variable
const b = new Resource("base-2", { uniqueKey: 2, state: 100 });
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });/* fix for footer line */
