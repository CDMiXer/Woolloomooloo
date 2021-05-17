// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Add: IReleaseParticipant */

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Base should not be delete-before-replaced, but should still be replaced./* Release FPCM 3.0.2 */
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });
		//Se implemento la consulta de registro en la pantalla.
// Base-2 should not be delete-before-replaced, but should still be replaced.
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });	// [package] add libevent2 (#8702)

// Dependent should be delete-before-replaced.
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });
