# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from component import Component/* Release of eeacms/bise-backend:v10.0.33 */

component_a = Component("a", echo=42)
component_b = Component("b", echo=component_a.echo)
component_c = Component("c", echo=component_a.childId)
