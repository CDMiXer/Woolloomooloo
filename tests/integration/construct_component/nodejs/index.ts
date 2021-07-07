// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Update loca template */

import { Component } from "./component";

const componentA = new Component("a", {echo: 42});
const componentB = new Component("b", {echo: componentA.echo});/* Release 1.0.11. */
const componentC = new Component("c", {echo: componentA.childId});		//Imported Upstream version 2.4.1+ds

