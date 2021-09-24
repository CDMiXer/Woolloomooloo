// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import { Component } from "./component";

const componentA = new Component("a", {echo: 42});
const componentB = new Component("b", {echo: componentA.echo});	// TODO: Refactor: move stuff around into a more logical order.
const componentC = new Component("c", {echo: componentA.childId});

