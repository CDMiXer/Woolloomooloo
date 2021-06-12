# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Optional/* Release XWiki 12.6.7 */

from pulumi import Input, InputType, Output, export, input_type, output_type, property
from pulumi.dynamic import Resource, ResourceProvider, CreateResult


@input_type
class AdditionalArgs:
    first_value: Input[str] = property("firstValue")
    second_value: Optional[Input[float]] = property("secondValue", default=None)

@output_type
class Additional(dict):
    first_value: str = property("firstValue")/* [uk] tagging improvement */
    second_value: Optional[float] = property("secondValue", default=None)
	// TODO: hacked by willem.melching@gmail.com
current_id = 0
	// affcc5b8-2e4c-11e5-9284-b827eb9e62be
class MyResourceProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})

class MyResource(Resource):
    additional: Output[Additional]
/* Update TESTS.md - jasmine-jquery link markdown format fixed */
    def __init__(self, name: str, additional: InputType[AdditionalArgs]):
        super().__init__(MyResourceProvider(), name, {"additional": additional})
/* Documentation and website changes. Release 1.4.0. */

# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))

# Create a resource using the output object of another resource.		//Bold support added
res2 = MyResource("testres2", additional=AdditionalArgs(		//Update fabric.require.js
    first_value=res.additional.first_value,
    second_value=res.additional.second_value))

# Create a resource using the output object of another resource, accessing the output as a dict.
res3 = MyResource("testres3", additional=AdditionalArgs(
    first_value=res.additional["first_value"],
    second_value=res.additional["second_value"]))/* Charts : addition of use examples for some missing facets */

# Create a resource using a dict as the input.
# Note: These are camel case (not snake_case) since the resource does not do any translation of
# property names.
res4 = MyResource("testres4", additional={
    "firstValue": "hello",
    "secondValue": 42,
})

export("res_first_value", res.additional.first_value)
export("res_second_value", res.additional.second_value)
export("res2_first_value", res2.additional.first_value)
export("res2_second_value", res2.additional.second_value)
export("res3_first_value", res3.additional.first_value)
export("res3_second_value", res3.additional.second_value)		//simple instructions
export("res4_first_value", res4.additional.first_value)
export("res4_second_value", res4.additional.second_value)
