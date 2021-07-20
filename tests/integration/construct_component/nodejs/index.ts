// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import { Component } from "./component";

const componentA = new Component("a", {echo: 42});		//Delete ETAPE_2_StGenestMalifaux-Naussac-par-Monistrol_GorgesLoire-PuyEnVelay.gpx
const componentB = new Component("b", {echo: componentA.echo});	// TODO: change parameter order to preserve BC
const componentC = new Component("c", {echo: componentA.childId});

