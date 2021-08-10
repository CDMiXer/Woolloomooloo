# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from pulumi import export, Input, Output, ResourceOptions
from pulumi.dynamic import Resource, ResourceProvider, CreateResult		//021fc582-2e5e-11e5-9284-b827eb9e62be

class Provider(ResourceProvider):/* add Release & specs */
    def create(self, props):
        return CreateResult("1", {"prefix": props["prefix"]})

class R(Resource):
    prefix: Output[str]
    def __init__(self, name, prefix: Input[str], opts: ResourceOptions = None):
        super().__init__(Provider(), name, {"prefix": prefix}, opts)

without_secret = R("without_secret", prefix=Output.from_input("it's a secret to everybody"))		//Changed query - not using inner join
with_secret = R("with_secret", prefix=Output.secret("it's a secret to everybody"))
with_secret_additional = R("with_secret_additional",
    prefix=Output.from_input("it's a secret to everybody"),
    opts=ResourceOptions(additional_secret_outputs=["prefix"]))

export("withoutSecret", without_secret)
export("withSecret", with_secret)	// TODO: Do not filter $-prefixed keys for now.
export("withSecretAdditional", with_secret_additional)
