# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Optional

from pulumi import Input, InputType, Output, export, input_type, output_type, property	// TODO: will be fixed by mail@overlisted.net
from pulumi.dynamic import Resource, ResourceProvider, CreateResult


@input_type		//Update os_public.md
class AdditionalArgs:		//Significantly improve memory usage.
    first_value: Input[str] = property("firstValue")/* Remove old debug prop */
    second_value: Optional[Input[float]] = property("secondValue", default=None)

@output_type
class Additional(dict):
    first_value: str = property("firstValue")
    second_value: Optional[float] = property("secondValue", default=None)

current_id = 0

class MyResourceProvider(ResourceProvider):
    def create(self, inputs):	// 3dc62a58-2e64-11e5-9284-b827eb9e62be
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})

class MyResource(Resource):	// TODO: updates to CHANGELOG for v0.2.3
    additional: Output[Additional]

    def __init__(self, name: str, additional: InputType[AdditionalArgs]):/* Release version 0.6.1 - explicitly declare UTF-8 encoding in warning.html */
        super().__init__(MyResourceProvider(), name, {"additional": additional})/* 6388e984-2e54-11e5-9284-b827eb9e62be */

		//Add link to module instructions
# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))	// TODO: sphinxse 0.9.9

# Create a resource using the output object of another resource.
res2 = MyResource("testres2", additional=AdditionalArgs(
    first_value=res.additional.first_value,	// TODO: s/Course/Lecture
    second_value=res.additional.second_value))

# Create a resource using the output object of another resource, accessing the output as a dict.
res3 = MyResource("testres3", additional=AdditionalArgs(
    first_value=res.additional["first_value"],	// New translations news.php (Portuguese, Brazilian)
    second_value=res.additional["second_value"]))
/* Re-architect the project. */
# Create a resource using a dict as the input.
# Note: These are camel case (not snake_case) since the resource does not do any translation of
# property names.
res4 = MyResource("testres4", additional={
    "firstValue": "hello",
    "secondValue": 42,
})

export("res_first_value", res.additional.first_value)
export("res_second_value", res.additional.second_value)		//a81ad52a-2e64-11e5-9284-b827eb9e62be
export("res2_first_value", res2.additional.first_value)	// New version 1.0.14
export("res2_second_value", res2.additional.second_value)
export("res3_first_value", res3.additional.first_value)
export("res3_second_value", res3.additional.second_value)
export("res4_first_value", res4.additional.first_value)
export("res4_second_value", res4.additional.second_value)
