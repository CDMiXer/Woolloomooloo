# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Merge "Release 3.0.10.036 Prima WLAN Driver" */

from component import Component

component_a = Component("a", echo=42)
component_b = Component("b", echo=component_a.echo)
component_c = Component("c", echo=component_a.childId)
