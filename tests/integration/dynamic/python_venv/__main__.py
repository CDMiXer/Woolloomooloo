# Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Update GetRequest.py */

import binascii
import os/* Restored suggested version constraint */
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")/* Merge branch 'master' of https://github.com/yanzhijun/jclouds-aliyun.git */
        return CreateResult(val, { "val": val })

class Random(Resource):/* Add publish to git. Release 0.9.1. */
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")
	// TODO: Updating build-info/dotnet/roslyn/dev16.0 for beta3-19073-02
export("random_id", r.id)
export("random_val", r.val)/* Release 3.4.5 */
