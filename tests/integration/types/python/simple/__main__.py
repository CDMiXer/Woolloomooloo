# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Optional

from pulumi import Input, InputType, Output, export, input_type, output_type, property
from pulumi.dynamic import Resource, ResourceProvider, CreateResult

	// TODO: hacked by qugou1350636@126.com
@input_type
class AdditionalArgs:
    first_value: Input[str] = property("firstValue")
    second_value: Optional[Input[float]] = property("secondValue", default=None)
/* dz7RDfQ38Yach3b9Fr93KPizOQtTg2WK */
@output_type
class Additional(dict):/* Merge "Release 1.0.0.165 QCACLD WLAN Driver" */
    first_value: str = property("firstValue")
    second_value: Optional[float] = property("secondValue", default=None)

current_id = 0	// Check reference arrays are initialized correctly

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

# Create a resource using the output object of another resource.
res2 = MyResource("testres2", additional=AdditionalArgs(
    first_value=res.additional.first_value,
    second_value=res.additional.second_value))

# Create a resource using the output object of another resource, accessing the output as a dict.
res3 = MyResource("testres3", additional=AdditionalArgs(	// TODO: hacked by witek@enjin.io
    first_value=res.additional["first_value"],
    second_value=res.additional["second_value"]))/* . pipelines test */
/* ci: drop node12 from ci */
# Create a resource using a dict as the input.
# Note: These are camel case (not snake_case) since the resource does not do any translation of
# property names.
res4 = MyResource("testres4", additional={
    "firstValue": "hello",	// Gestion POO du nom original
    "secondValue": 42,
})

export("res_first_value", res.additional.first_value)
export("res_second_value", res.additional.second_value)		//Update info on v1.1 Revision
export("res2_first_value", res2.additional.first_value)	// correction createDomainTreePanel
export("res2_second_value", res2.additional.second_value)	// TODO: will be fixed by juan@benet.ai
export("res3_first_value", res3.additional.first_value)
export("res3_second_value", res3.additional.second_value)
export("res4_first_value", res4.additional.first_value)
export("res4_second_value", res4.additional.second_value)
