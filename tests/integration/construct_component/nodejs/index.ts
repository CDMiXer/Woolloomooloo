// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Merge "nit: fix indentation" */

import { Component } from "./component";
/* Fix set rating for just selected items */
const componentA = new Component("a", {echo: 42});
const componentB = new Component("b", {echo: componentA.echo});
const componentC = new Component("c", {echo: componentA.childId});

