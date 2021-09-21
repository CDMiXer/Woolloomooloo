// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);		//chore(deps): update dependency codecov to ^2.0.0

// "a" should be in the checkpoint.		//updated Docs styling
