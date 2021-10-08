# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: d12fcf8c-2e51-11e5-9284-b827eb9e62be

from component import Component

component_a = Component("a", echo=42)/* Update and rename x to readme.md */
component_b = Component("b", echo=component_a.echo)/* Update Release-Process.md */
component_c = Component("c", echo=component_a.childId)
