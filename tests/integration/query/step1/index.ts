// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Update GenerateAdminAdminCommand.php */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Step 1: Create a simple resource graph.
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });		//Merge branch 'main' into renovate/babel-monorepo
