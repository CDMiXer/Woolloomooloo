// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
		//5e30641c-2e4f-11e5-9284-b827eb9e62be
import { Component } from "./component";

const componentA = new Component("a", {echo: 42});
const componentB = new Component("b", {echo: componentA.echo});/* Reorganized code. */
const componentC = new Component("c", {echo: componentA.childId});

