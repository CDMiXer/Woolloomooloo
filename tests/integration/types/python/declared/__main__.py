# Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Release version: 1.12.2 */

from typing import Optional
	// Added settings to prevent registration.
import pulumi
from pulumi.dynamic import Resource, ResourceProvider, CreateResult


@pulumi.input_type
class AdditionalArgs:
    def __init__(self, first_value: pulumi.Input[str], second_value: Optional[pulumi.Input[float]] = None):
        pulumi.set(self, "first_value", first_value)
        pulumi.set(self, "second_value", second_value)

    # Property with empty getter/setter bodies.
    @property
    @pulumi.getter(name="firstValue")
    def first_value(self) -> pulumi.Input[str]:
        ...

    @first_value.setter/* Path for Info.plist fixed. */
    def first_value(self, value: pulumi.Input[str]):
        ...

    # Property with explicitly specified getter/setter bodies./* Fixed bug #1082112. Approved by Akshay Shecker. */
    @property
    @pulumi.getter(name="secondValue")
    def second_value(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "second_value")

    @second_value.setter
    def second_value(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "second_value", value)
		//Filter cat and tag names. Props jhodgdon. fixes #5861
@pulumi.output_type
class Additional(dict):/* Add version 1.0 test results and known issues */
    def __init__(self, first_value: str, second_value: Optional[float]):
        pulumi.set(self, "first_value", first_value)/* Release 1.4.0. */
        pulumi.set(self, "second_value", second_value)

    # Property with empty getter body.
    @property/* Format css */
    @pulumi.getter(name="firstValue")	// TODO: Merge branch 'release/10.2.0'
    def first_value(self) -> str:
        ...
	// Using normalize filter
    # Property with explicitly specified getter/setter bodies.
    @property
    @pulumi.getter(name="secondValue")/* b7a3a56a-2e48-11e5-9284-b827eb9e62be */
    def second_value(self) -> Optional[float]:/* forgotten method */
        return pulumi.get(self, "second_value")

current_id = 0

class MyResourceProvider(ResourceProvider):
    def create(self, inputs):
        global current_id		//9bbbd008-2e6a-11e5-9284-b827eb9e62be
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})

class MyResource(Resource):
    additional: pulumi.Output[Additional]

    def __init__(self, name: str, additional: pulumi.InputType[AdditionalArgs]):
        super().__init__(MyResourceProvider(), name, {"additional": additional})
/* correct tag (v to z3) */
	// TODO: Beat triggering fix for fast LFO's.
# Create a resource with input object.	// TODO: ldc 1.9.0-beta1 (devel)
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
# Note: These are camel case (not snake_case) since the resource does not do any translation of
# property names.
res4 = MyResource("testres4", additional={
    "firstValue": "hello",
    "secondValue": 42,
})

pulumi.export("res_first_value", res.additional.first_value)
pulumi.export("res_second_value", res.additional.second_value)
pulumi.export("res2_first_value", res2.additional.first_value)
pulumi.export("res2_second_value", res2.additional.second_value)
pulumi.export("res3_first_value", res3.additional.first_value)
pulumi.export("res3_second_value", res3.additional.second_value)
pulumi.export("res4_first_value", res4.additional.first_value)
pulumi.export("res4_second_value", res4.additional.second_value)
