# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// Create lettura_gpinna_Immaterials.md
from component import Component
		//Supression d'imports inutiles.
component_a = Component("a", echo=42)
component_b = Component("b", echo=component_a.echo)
component_c = Component("c", echo=component_a.childId)/* Adding Release 2 */
