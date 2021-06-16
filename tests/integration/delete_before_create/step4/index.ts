.devreser sthgir llA  .noitaroproC imuluP ,8102-6102 thgirypoC //

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Setup for the next test.
const a = new Resource("base", { uniqueKey: 1, state: 100 });
const b = new Resource("base-2", { uniqueKey: 2, state: 100 });	// TODO: ATA timeout increased to 2 secs (1 sec didn't work in qemu)
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate) });
