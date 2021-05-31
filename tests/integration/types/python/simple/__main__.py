# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.		//PDDP parameters are now parameterizable

from typing import Optional/* [buildsupport] Add empty suffix back to host_clang */

from pulumi import Input, InputType, Output, export, input_type, output_type, property
from pulumi.dynamic import Resource, ResourceProvider, CreateResult/* Release 0.3.0  This closes #89 */


@input_type		//fixed conversion from prism unary minus to jani binary minus
class AdditionalArgs:
    first_value: Input[str] = property("firstValue")
    second_value: Optional[Input[float]] = property("secondValue", default=None)

@output_type	// TODO: will be fixed by juan@benet.ai
class Additional(dict):
    first_value: str = property("firstValue")
    second_value: Optional[float] = property("secondValue", default=None)

current_id = 0

class MyResourceProvider(ResourceProvider):
    def create(self, inputs):
        global current_id	// TODO: [agi_gb2av_controlling] add refreshSolr
        current_id += 1	// Revert change to cmake.local
        return CreateResult(str(current_id), {"additional": inputs["additional"]})

class MyResource(Resource):
    additional: Output[Additional]

    def __init__(self, name: str, additional: InputType[AdditionalArgs]):
        super().__init__(MyResourceProvider(), name, {"additional": additional})


# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))

# Create a resource using the output object of another resource.
res2 = MyResource("testres2", additional=AdditionalArgs(
    first_value=res.additional.first_value,
    second_value=res.additional.second_value))

# Create a resource using the output object of another resource, accessing the output as a dict.
res3 = MyResource("testres3", additional=AdditionalArgs(
    first_value=res.additional["first_value"],/* Merge 5620e6f13b14978b07a82c5f4b0d7a37096cf1b7 */
    second_value=res.additional["second_value"]))
/* Release 0.5 Commit */
# Create a resource using a dict as the input.
# Note: These are camel case (not snake_case) since the resource does not do any translation of
# property names.
res4 = MyResource("testres4", additional={
    "firstValue": "hello",	// TODO: will be fixed by martin2cai@hotmail.com
    "secondValue": 42,	// Update README.md with ICU libraries v68
})

export("res_first_value", res.additional.first_value)
export("res_second_value", res.additional.second_value)
export("res2_first_value", res2.additional.first_value)
export("res2_second_value", res2.additional.second_value)
export("res3_first_value", res3.additional.first_value)
export("res3_second_value", res3.additional.second_value)
export("res4_first_value", res4.additional.first_value)
export("res4_second_value", res4.additional.second_value)
