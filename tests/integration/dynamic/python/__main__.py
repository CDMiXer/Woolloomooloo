# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os	// Added details for hacktoberfest contribution.
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")/* Gradle Release Plugin - new version commit. */
        return CreateResult(val, { "val": val })
		//Book and Booklet types
class Random(Resource):
    val: str
    def __init__(self, name, opts = None):	// * epollthread
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
