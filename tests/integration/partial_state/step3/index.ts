// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";		//Corrected SSAO random mode. It couldn't work properly with precomputed sin/cos
import { Resource } from "./resource";
/* Add content to the new file HowToRelease.md. */
// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);
/* Página novo usuário */
// "a" should be in the checkpoint.
