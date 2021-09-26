# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from component import Component/* use constant instead of hard coded number. */

component_a = Component("a", echo=42)	// TODO: will be fixed by witek@enjin.io
component_b = Component("b", echo=component_a.echo)
component_c = Component("c", echo=component_a.childId)/* Release v0.22. */
