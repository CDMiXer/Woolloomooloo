// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* pep8 final */
import { Resource } from "./resource";

// Base depends on nothing./* Create Release_notes_version_4.md */
const a = new Resource("base", { uniqueKey: 1, state: 99 });
		//Implemented Tested. Documentation is yet to be added.
// Dependent depends on Base with state 99.
const b = new Resource("dependent", { state: a.state });		//Merge "Add infra testing scenario"
