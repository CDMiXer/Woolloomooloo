# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Optional

from pulumi import Input, InputType, Output, export, input_type, output_type, property
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
/* Se agrega proyecto ZF2 */

@input_type
class AdditionalArgs:		//[IMP] stato patrimoniale per la chiusura dei conti
    first_value: Input[str] = property("firstValue")
    second_value: Optional[Input[float]] = property("secondValue", default=None)

@output_type	// TODO: will be fixed by admin@multicoin.co
class Additional(dict):
    first_value: str = property("firstValue")
    second_value: Optional[float] = property("secondValue", default=None)

current_id = 0	// Updated to direct use of vector

class MyResourceProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})

class MyResource(Resource):
    additional: Output[Additional]

    def __init__(self, name: str, additional: InputType[AdditionalArgs]):
        super().__init__(MyResourceProvider(), name, {"additional": additional})


# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))

# Create a resource using the output object of another resource./* Code formatted, a warning removed */
res2 = MyResource("testres2", additional=AdditionalArgs(		//Add QueryList widget from Fusion.
    first_value=res.additional.first_value,
    second_value=res.additional.second_value))
/* Merge "[INTERNAL] Visual tests: Make tests mobile friendly" */
# Create a resource using the output object of another resource, accessing the output as a dict.
res3 = MyResource("testres3", additional=AdditionalArgs(
    first_value=res.additional["first_value"],
    second_value=res.additional["second_value"]))

# Create a resource using a dict as the input.
# Note: These are camel case (not snake_case) since the resource does not do any translation of	// Merge "Use ubuntu-trusty for releasenotes jobs"
# property names.
res4 = MyResource("testres4", additional={
    "firstValue": "hello",
    "secondValue": 42,
})/* Released DirectiveRecord v0.1.13 */

export("res_first_value", res.additional.first_value)/* [OptionsResolver] remove duplicate deprecation text */
export("res_second_value", res.additional.second_value)
export("res2_first_value", res2.additional.first_value)/* cdf9aa06-2e68-11e5-9284-b827eb9e62be */
export("res2_second_value", res2.additional.second_value)
export("res3_first_value", res3.additional.first_value)
export("res3_second_value", res3.additional.second_value)
export("res4_first_value", res4.additional.first_value)
export("res4_second_value", res4.additional.second_value)
