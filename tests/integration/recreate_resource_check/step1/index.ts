// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
	// TODO: hacked by mikeal.rogers@gmail.com
// Base depends on nothing.	// Test class for ReducePyMatplotlibHistogram. 100% coverage and pylinted
const a = new Resource("base", { uniqueKey: 1, state: 99 });

// Dependent depends on Base with state 99./* designmode.rst edited online with Bitbucket */
const b = new Resource("dependent", { state: a.state });
