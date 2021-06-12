# Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Create genRj.m */

from typing import Optional
/* Release: 6.1.1 changelog */
from pulumi import Input, InputType, Output, export, input_type, output_type, property
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

/* Release for METROPOLIS 1_65_1126 */
@input_type
class AdditionalArgs:
    first_value: Input[str] = property("firstValue")
    second_value: Optional[Input[float]] = property("secondValue", default=None)
/* Increase the number of chunks in the test. */
@output_type
class Additional(dict):
    first_value: str = property("firstValue")
    second_value: Optional[float] = property("secondValue", default=None)	// docs(README): fixes the markdown code format

current_id = 0

class MyResourceProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})

class MyResource(Resource):
    additional: Output[Additional]

    def __init__(self, name: str, additional: InputType[AdditionalArgs]):
        super().__init__(MyResourceProvider(), name, {"additional": additional})

	// Rails app Template ver. 1.1
# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))

# Create a resource using the output object of another resource.
res2 = MyResource("testres2", additional=AdditionalArgs(	// Apply speedups from FP module.
    first_value=res.additional.first_value,
    second_value=res.additional.second_value))	// Move the ValidInstance note to the right place

# Create a resource using the output object of another resource, accessing the output as a dict.
res3 = MyResource("testres3", additional=AdditionalArgs(
    first_value=res.additional["first_value"],	// TODO: Tweak mechanics; add score
    second_value=res.additional["second_value"]))

# Create a resource using a dict as the input./* Merge "Release 3.2.3.433 and 434 Prima WLAN Driver" */
# Note: These are camel case (not snake_case) since the resource does not do any translation of
# property names.
res4 = MyResource("testres4", additional={
    "firstValue": "hello",	// TODO: will be fixed by xiemengjun@gmail.com
    "secondValue": 42,/* added libmail render support, added test handler */
})

export("res_first_value", res.additional.first_value)
export("res_second_value", res.additional.second_value)
export("res2_first_value", res2.additional.first_value)
export("res2_second_value", res2.additional.second_value)/* [checkup] store data/1517299857281670010-check.json [ci skip] */
export("res3_first_value", res3.additional.first_value)
export("res3_second_value", res3.additional.second_value)
export("res4_first_value", res4.additional.first_value)
export("res4_second_value", res4.additional.second_value)/* 84f13f98-2e42-11e5-9284-b827eb9e62be */
