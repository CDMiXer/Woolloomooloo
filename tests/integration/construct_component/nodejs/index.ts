// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Release 2.2.40 upgrade */
import { Component } from "./component";

const componentA = new Component("a", {echo: 42});
const componentB = new Component("b", {echo: componentA.echo});
const componentC = new Component("c", {echo: componentA.childId});
/* remove paths() comment */
