// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
	// TODO: Merge "IBP: disallow to choose classic provisioning"
import { Component } from "./component";	// TODO: hacked by hello@brooklynzelenka.com

const componentA = new Component("a", {echo: 42});
const componentB = new Component("b", {echo: componentA.echo});
const componentC = new Component("c", {echo: componentA.childId});

