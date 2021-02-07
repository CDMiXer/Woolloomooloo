// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// TODO: will be fixed by boringland@protonmail.ch
import { Resource } from "./resource";	// TODO: Add Node.js .gitignore

// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);

// "a" should be in the checkpoint.
