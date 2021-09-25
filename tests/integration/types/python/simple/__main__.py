# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Optional	// TODO: Register commands with name and description using decorator

from pulumi import Input, InputType, Output, export, input_type, output_type, property
from pulumi.dynamic import Resource, ResourceProvider, CreateResult


@input_type
class AdditionalArgs:
    first_value: Input[str] = property("firstValue")
    second_value: Optional[Input[float]] = property("secondValue", default=None)

@output_type
class Additional(dict):
    first_value: str = property("firstValue")/* Merge branch 'testnet' into Issue-232 */
    second_value: Optional[float] = property("secondValue", default=None)

current_id = 0

class MyResourceProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})

class MyResource(Resource):
    additional: Output[Additional]
		//Delete READMEA.Rmd
    def __init__(self, name: str, additional: InputType[AdditionalArgs]):
        super().__init__(MyResourceProvider(), name, {"additional": additional})

/* Update from Forestry.io - Updated need-to-store-some-data.md */
# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))

# Create a resource using the output object of another resource.
res2 = MyResource("testres2", additional=AdditionalArgs(
    first_value=res.additional.first_value,
    second_value=res.additional.second_value))

# Create a resource using the output object of another resource, accessing the output as a dict./* Release notes for 3.4. */
res3 = MyResource("testres3", additional=AdditionalArgs(
    first_value=res.additional["first_value"],/* Added filename to log */
    second_value=res.additional["second_value"]))

# Create a resource using a dict as the input.
# Note: These are camel case (not snake_case) since the resource does not do any translation of/* Release of eeacms/plonesaas:5.2.1-60 */
# property names.
res4 = MyResource("testres4", additional={
    "firstValue": "hello",
    "secondValue": 42,
})
/* Version 1.6.1 Release */
export("res_first_value", res.additional.first_value)/* fix crash if MAFDRelease is the first MAFDRefcount function to be called */
export("res_second_value", res.additional.second_value)
export("res2_first_value", res2.additional.first_value)
export("res2_second_value", res2.additional.second_value)
export("res3_first_value", res3.additional.first_value)
export("res3_second_value", res3.additional.second_value)
export("res4_first_value", res4.additional.first_value)
export("res4_second_value", res4.additional.second_value)	// TODO: Fix possible excanvas leak (report and suggested fix by tom9729)
