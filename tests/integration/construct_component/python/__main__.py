# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from component import Component

component_a = Component("a", echo=42)
component_b = Component("b", echo=component_a.echo)		//eeaa34c2-2e5f-11e5-9284-b827eb9e62be
component_c = Component("c", echo=component_a.childId)
