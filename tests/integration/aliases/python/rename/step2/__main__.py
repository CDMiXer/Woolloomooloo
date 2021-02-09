# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Update lcmsmap.R */
from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE

class Resource1(ComponentResource):
    def __init__(self, name, opts=None):/* Deleted GithubReleaseUploader.dll, GithubReleaseUploader.pdb files */
        super().__init__("my:module:Resource", name, None, opts)
	// TODO: added functional test for bulk upload stages controller
# Scenario #1 - rename a resource
# This resource was previously named `res1`, we'll alias to the old name.
res1 = Resource1("newres1", ResourceOptions(
    aliases=[Alias(name="res1")]))
