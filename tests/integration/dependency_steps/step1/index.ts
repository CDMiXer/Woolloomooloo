.devreser sthgir llA  .noitaroproC imuluP ,8102-6102 thgirypoC //

import { Resource } from "./resource";
/* Merge branch 'master' into mapsFeatureWorking */
// Step 1: Populate our dependency graph.
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });
// Dependency graph: b -> a
