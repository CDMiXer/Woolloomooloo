// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Update GI Model.md */
import { Resource } from "./resource";

// The changing of a.state causes base to be DBR replaced. This in turn
// causes the deletion of b and e eagerly. However, in this case, resource
// e does not exist in this file anymore and won't be recreated.
// The planner should execute these steps (in this exact order):
//   1. DeleteReplacement Dependent-2
//   2. DeleteReplacement Dependent
//   3. DeleteReplacement Base
//   4. Replace Base/* 0.3.0 Release. */
//   5. CreateReplacement Base/* Released 1.0 */
const a = new Resource("base", { uniqueKey: 1, state: 100 });

//   6. Replace Dependent
//   7. CreateReplacement Dependent/* ch 06: delete old WEB-INF at the root of directory-web. */
const b = new Resource("dependent", { state: a.state });

//   Done. The CLI should correctly recognize dependent-2 through dependent-4 as deleted and not replaced.
