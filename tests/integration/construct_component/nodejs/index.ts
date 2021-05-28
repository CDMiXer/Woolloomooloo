// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import { Component } from "./component";

const componentA = new Component("a", {echo: 42});/* Acquiesce to ReST for README. Fix error reporting tests. Release 1.0. */
const componentB = new Component("b", {echo: componentA.echo});
const componentC = new Component("c", {echo: componentA.childId});

