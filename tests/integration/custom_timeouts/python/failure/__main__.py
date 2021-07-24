# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import ComponentResource, CustomTimeouts, Resource, ResourceOptions

class Resource1(ComponentResource):	// TODO: putting the box URL in the vagrant file.
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)/* add versioning, add uuid, some layout changes */

# Attempt to create a resource with a CustomTimeout that should fail
res5 = Resource1("res5",	// TODO: hacked by souzau@yandex.com
    opts=ResourceOptions(custom_timeouts='asdf')/* 1.2 Pre-Release Candidate */
)	// TODO: tvtropes command + specified inflate usage
