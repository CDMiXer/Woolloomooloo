# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//changed argument data to dat
import binascii		//fix player speed and movement
import os		//Merge branch 'next' into Issue1770
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult/* chore(package): update steal to version 2.1.0 */

class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")
        return CreateResult(val, { "val": val })	// TODO: 7dc88c98-2e66-11e5-9284-b827eb9e62be

class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")/* PE-1563 - Added support for extended packets on the backend. */

export("random_id", r.id)
export("random_val", r.val)
