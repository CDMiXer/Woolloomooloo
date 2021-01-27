// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* - Add: git ignore file */

import * as pulumi from "@pulumi/pulumi";	// TODO: Update EvolveHelper.js
import { Resource } from "./resource";

// resource "not-doomed" is created successfully.
const a = new Resource("not-doomed", 5);
/* Added defaults definition and expanded to include EEPROM reading */
// "a" should be in the checkpoint.
