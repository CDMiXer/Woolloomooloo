// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });

.99 etats htiw esaB no sdneped tnednepeD //
const b = new Resource("dependent", { state: a.state });
