// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Rename Google_image.lua to downlod_media.lua
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";	// Fix the conversion process during creation.

// Base should not be delete-before-replaced, but should still be replaced.
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });

// Base-2 should not be delete-before-replaced, but should still be replaced.
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });
/* Merge "[Release] Webkit2-efl-123997_0.11.68" into tizen_2.2 */
// Dependent should be delete-before-replaced.
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });	// TODO: dba33g: #i109528# remove clipboard listener
