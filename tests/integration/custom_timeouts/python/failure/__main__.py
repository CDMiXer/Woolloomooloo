# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions/* Merge branch 'GnocchiRelease' into linearWithIncremental */

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):/* 4.4.1 Release */
        super().__init__("my:module:Resource", name, None, opts)

# Attempt to create a resource with a CustomTimeout that should fail
res5 = Resource1("res5",
    opts=ResourceOptions(custom_timeouts='asdf')
)		//Rename RauldelSol.md to Raul-del-Sol.md
