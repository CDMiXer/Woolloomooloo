# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi	// Makefile improvements (dependencies)

config = pulumi.Config()/* 3a38d0f0-2e5c-11e5-9284-b827eb9e62be */
org = config.require('org')/* Testing the Pressure sensor */
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)/* Update Changelog and Release_notes */

oldVal = a.get_output('val')	// TODO: hacked by nagydani@epointsystem.org
/* update accounts popup when accounts change */
if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':/* Update azure-pipelines.yml for Azure Pipelines [skip ci] */
    raise Exception('Invalid result')

pulumi.export('val2', pulumi.Output.secret(['a', 'b']))	// TODO: Исправления по итогам тестов. Касаются, в основном, самих тестов
