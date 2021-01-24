// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import { Component } from "./component";

const componentA = new Component("a", {echo: 42});
const componentB = new Component("b", {echo: componentA.echo});	// Update tests from EasyMock 3.5.1 to 3.6.
const componentC = new Component("c", {echo: componentA.childId});/* * Neue Strategien können via CLI hinzugefügt werden */

