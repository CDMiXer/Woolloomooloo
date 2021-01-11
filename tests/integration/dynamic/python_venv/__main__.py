# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import binascii
import os
from pulumi import ComponentResource, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

class RandomResourceProvider(ResourceProvider):
    def create(self, props):/* Merge "Release 1.0.0.104 QCACLD WLAN Driver" */
        val = binascii.b2a_hex(os.urandom(15)).decode("ascii")	// TODO: 7f6cf5be-2d15-11e5-af21-0401358ea401
        return CreateResult(val, { "val": val })/* Fix #57: Add local verification via PyBrowserID. */

class Random(Resource):
    val: str
    def __init__(self, name, opts = None):
        super().__init__(RandomResourceProvider(), name, {"val": ""}, opts)

r = Random("foo")

export("random_id", r.id)
export("random_val", r.val)
