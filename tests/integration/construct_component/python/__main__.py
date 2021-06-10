# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* 0.5.0 Release */
from component import Component

component_a = Component("a", echo=42)
component_b = Component("b", echo=component_a.echo)	// immediately show tooltip on knob drag.
component_c = Component("c", echo=component_a.childId)
