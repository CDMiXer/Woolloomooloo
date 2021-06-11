// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Spark Environment File.

import { Resource } from "./resource";

// Base changes its state to 21, triggering DBR replacement.
const a = new Resource("base", { uniqueKey: 1, state: 21 });/* [releng] Release Snow Owl v6.16.3 */

// The DBR replacement of Base triggers an early deletion of dependent.	// TODO: Add docstrings for BrokerConfigurationHelper and ExchangeHelper

.22 etats htiw ereh tnedneped etaerc-er lliw enigne eht ,esab fo noitaerc-er eht retfA //
// The engine should not consider the old state of "dependent" (namely 99) when running	// TODO: hacked by peterke@gmail.com
// Check on this new resource with state 22.
const b = new Resource("dependent", { state: a.state.apply((num: number) => num + 1) });
