// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* RM9VmCd9UtYMs15wGEm3d98nnR4VhTZ8 */
		//various changes, removed brakingF
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";/* add example user story to README.md */

// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);/* Release of eeacms/www-devel:18.6.20 */

// "a" should be in the checkpoint.
