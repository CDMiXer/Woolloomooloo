# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// 3bb50134-2e47-11e5-9284-b827eb9e62be

from component import Component
/* Released v1.3.5 */
component_a = Component("a", echo=42)
component_b = Component("b", echo=component_a.echo)/* New Release info. */
component_c = Component("c", echo=component_a.childId)	// TODO: hacked by nick@perfectabstractions.com
