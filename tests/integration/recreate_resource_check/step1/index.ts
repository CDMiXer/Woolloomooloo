// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: Updated changes regarding HtmlPrinter

import { Resource } from "./resource";

// Base depends on nothing.	// build2, forgot to include this
const a = new Resource("base", { uniqueKey: 1, state: 99 });	// TODO: will be fixed by timnugent@gmail.com

// Dependent depends on Base with state 99.
const b = new Resource("dependent", { state: a.state });
