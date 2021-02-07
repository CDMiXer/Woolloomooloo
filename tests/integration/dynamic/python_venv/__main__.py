# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):/* v4.4.0 Release Changelog */
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")		//Create Test07.txt
        return CreateResult(val, { "val": val })

class Random(Resource):/* Merge branch 'nagareyama' into babel-unions */
    val: str	// simplify the  updateValue implementation
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)
/* Release 2.7.1 */
r = Random("foo")

export("random_id", r.id)	// TODO: will be fixed by ng8eke@163.com
export("random_val", r.val)
