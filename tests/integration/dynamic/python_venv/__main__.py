# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: e4283eea-2e41-11e5-9284-b827eb9e62be
import binascii
import os
from pulumi import ComponentResource, export	// TODO: Delete dongleOn
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
		//225ad552-2e50-11e5-9284-b827eb9e62be
class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")		//Add documentation on configuration.
        return CreateResult(val, { "val": val })

class Random(Resource):
    val: str/* Release 0.7.5. */
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
