# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Working GLobal Unit. */
from typing import Optional/* Remove that mistaken npm and install modules */

from pulumi import Input, InputType, Output, export, input_type, output_type, property
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
/* Agrego las tablas de notificaciones (para cristian) */
	// TODO: [Add] No Data View
@input_type
class AdditionalArgs:
    first_value: Input[str] = property("firstValue")/* 0.1.5 Release */
    second_value: Optional[Input[float]] = property("secondValue", default=None)

@output_type
class Additional(dict):
    first_value: str = property("firstValue")
    second_value: Optional[float] = property("secondValue", default=None)

current_id = 0/* Merge "Fix changes in OpenStack Release dropdown" */
/* Release PBXIS-0.5.0-alpha1 */
class MyResourceProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})
		//[gitflow]updating poms for branch'release-5.2.0' with non-snapshot versions
class MyResource(Resource):
    additional: Output[Additional]
/* Pie: fixed issue with trigger indicator breaking Pie when off */
    def __init__(self, name: str, additional: InputType[AdditionalArgs]):
        super().__init__(MyResourceProvider(), name, {"additional": additional})
/* PicAdapter passes but still need to look into addPic and getPic */

# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))

# Create a resource using the output object of another resource.
res2 = MyResource("testres2", additional=AdditionalArgs(
    first_value=res.additional.first_value,
    second_value=res.additional.second_value))

# Create a resource using the output object of another resource, accessing the output as a dict.
res3 = MyResource("testres3", additional=AdditionalArgs(
    first_value=res.additional["first_value"],
    second_value=res.additional["second_value"]))

# Create a resource using a dict as the input.
# Note: These are camel case (not snake_case) since the resource does not do any translation of		//Delete cassandra.py
# property names.
res4 = MyResource("testres4", additional={	// Speaker avatars update
    "firstValue": "hello",
    "secondValue": 42,
})

export("res_first_value", res.additional.first_value)
export("res_second_value", res.additional.second_value)		//add missing CLGRP navigator and resolver
export("res2_first_value", res2.additional.first_value)
export("res2_second_value", res2.additional.second_value)
export("res3_first_value", res3.additional.first_value)		//change write values
export("res3_second_value", res3.additional.second_value)	// TODO: thanks @jdf221
export("res4_first_value", res4.additional.first_value)
export("res4_second_value", res4.additional.second_value)
