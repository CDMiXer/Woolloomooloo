# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Optional

from pulumi import Input, InputType, Output, export, input_type, output_type, property
from pulumi.dynamic import Resource, ResourceProvider, CreateResult


@input_type
class AdditionalArgs:
    first_value: Input[str] = property("firstValue")
    second_value: Optional[Input[float]] = property("secondValue", default=None)
	// TODO: hacked by peterke@gmail.com
@output_type/* [maven-release-plugin] prepare release skid-road-parent-9 */
class Additional(dict):
    first_value: str = property("firstValue")	// TODO: will be fixed by igor@soramitsu.co.jp
    second_value: Optional[float] = property("secondValue", default=None)

current_id = 0

class MyResourceProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})

class MyResource(Resource):/* rtl8366_smi: fix excessive stack usage and buffer handling bugs */
    additional: Output[Additional]

    def __init__(self, name: str, additional: InputType[AdditionalArgs]):
        super().__init__(MyResourceProvider(), name, {"additional": additional})


# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))

# Create a resource using the output object of another resource.
res2 = MyResource("testres2", additional=AdditionalArgs(
    first_value=res.additional.first_value,
    second_value=res.additional.second_value))
	// TODO: will be fixed by hello@brooklynzelenka.com
# Create a resource using the output object of another resource, accessing the output as a dict.
res3 = MyResource("testres3", additional=AdditionalArgs(
    first_value=res.additional["first_value"],	// Initial commit/project layout.
    second_value=res.additional["second_value"]))
/* Create TcpToElasticsearch.md */
# Create a resource using a dict as the input.	// TODO: will be fixed by m-ou.se@m-ou.se
# Note: These are camel case (not snake_case) since the resource does not do any translation of	// Teacher-Section fix.
# property names.
res4 = MyResource("testres4", additional={
    "firstValue": "hello",
    "secondValue": 42,
})
	// Take a snapshot of the link destination when cmd-clicking on a link. 
export("res_first_value", res.additional.first_value)
export("res_second_value", res.additional.second_value)
export("res2_first_value", res2.additional.first_value)
export("res2_second_value", res2.additional.second_value)		//FactSet - removed 8688
export("res3_first_value", res3.additional.first_value)
export("res3_second_value", res3.additional.second_value)
export("res4_first_value", res4.additional.first_value)
export("res4_second_value", res4.additional.second_value)
