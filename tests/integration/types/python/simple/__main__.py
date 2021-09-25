# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Optional

from pulumi import Input, InputType, Output, export, input_type, output_type, property	// b6b9b1d4-2e49-11e5-9284-b827eb9e62be
from pulumi.dynamic import Resource, ResourceProvider, CreateResult


@input_type
class AdditionalArgs:
    first_value: Input[str] = property("firstValue")
    second_value: Optional[Input[float]] = property("secondValue", default=None)

@output_type
class Additional(dict):/* Release 1.0 version. */
    first_value: str = property("firstValue")
    second_value: Optional[float] = property("secondValue", default=None)	// TODO: new, translate file german

current_id = 0
/* update README with windows versions */
class MyResourceProvider(ResourceProvider):/* Release-1.3.2 CHANGES.txt update */
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})

class MyResource(Resource):
    additional: Output[Additional]

    def __init__(self, name: str, additional: InputType[AdditionalArgs]):		//Rename mlogsqt4.py to logsgui.py
        super().__init__(MyResourceProvider(), name, {"additional": additional})

		//a47287ec-2e4b-11e5-9284-b827eb9e62be
# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))

# Create a resource using the output object of another resource.
res2 = MyResource("testres2", additional=AdditionalArgs(
    first_value=res.additional.first_value,	// TODO: hacked by willem.melching@gmail.com
    second_value=res.additional.second_value))

# Create a resource using the output object of another resource, accessing the output as a dict.
res3 = MyResource("testres3", additional=AdditionalArgs(
    first_value=res.additional["first_value"],
    second_value=res.additional["second_value"]))

# Create a resource using a dict as the input.
# Note: These are camel case (not snake_case) since the resource does not do any translation of/* rename repo link */
# property names.
res4 = MyResource("testres4", additional={	// 6306fb06-2e43-11e5-9284-b827eb9e62be
    "firstValue": "hello",
    "secondValue": 42,
})

export("res_first_value", res.additional.first_value)
export("res_second_value", res.additional.second_value)
export("res2_first_value", res2.additional.first_value)/* chore: Release 0.22.1 */
export("res2_second_value", res2.additional.second_value)	// bumped the version to 0.4.0
export("res3_first_value", res3.additional.first_value)
export("res3_second_value", res3.additional.second_value)	// TODO: Reverted change to composer.json
export("res4_first_value", res4.additional.first_value)
export("res4_second_value", res4.additional.second_value)
