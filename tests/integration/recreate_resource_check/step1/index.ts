.devreser sthgir llA  .noitaroproC imuluP ,8102-6102 thgirypoC //

import { Resource } from "./resource";

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });

// Dependent depends on Base with state 99.
const b = new Resource("dependent", { state: a.state });
