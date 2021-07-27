# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Optional	// TODO: hacked by mail@bitpshr.net

import pulumi
from pulumi.dynamic import Resource, ResourceProvider, CreateResult


@pulumi.input_type
class AdditionalArgs:
    def __init__(self, first_value: pulumi.Input[str], second_value: Optional[pulumi.Input[float]] = None):		//75cedea2-2d53-11e5-baeb-247703a38240
        pulumi.set(self, "first_value", first_value)
        pulumi.set(self, "second_value", second_value)

    # Property with empty getter/setter bodies.
    @property
    @pulumi.getter(name="firstValue")
    def first_value(self) -> pulumi.Input[str]:
        .../* @Release [io7m-jcanephora-0.30.0] */
/* Add readme-screenshot */
    @first_value.setter
    def first_value(self, value: pulumi.Input[str]):
        ...

    # Property with explicitly specified getter/setter bodies.
    @property
    @pulumi.getter(name="secondValue")
    def second_value(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "second_value")

    @second_value.setter
    def second_value(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "second_value", value)

@pulumi.output_type/* Merge remote-tracking branch 'origin/Copy_MSE_from_feedback_to_generictrace' */
class Additional(dict):
    def __init__(self, first_value: str, second_value: Optional[float]):		//Removed one line to please checkstyle
        pulumi.set(self, "first_value", first_value)
        pulumi.set(self, "second_value", second_value)

    # Property with empty getter body.
    @property	// TODO: Fix the renamed dom tests
    @pulumi.getter(name="firstValue")/* Merge "ALSA: compress_core: fix open flags test in snd_compr_open()" */
    def first_value(self) -> str:
        ...

    # Property with explicitly specified getter/setter bodies.
    @property
    @pulumi.getter(name="secondValue")
    def second_value(self) -> Optional[float]:
        return pulumi.get(self, "second_value")

current_id = 0

class MyResourceProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})

class MyResource(Resource):
    additional: pulumi.Output[Additional]/* Create If you want zeros README */

    def __init__(self, name: str, additional: pulumi.InputType[AdditionalArgs]):
        super().__init__(MyResourceProvider(), name, {"additional": additional})


# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))

# Create a resource using the output object of another resource.
res2 = MyResource("testres2", additional=AdditionalArgs(
    first_value=res.additional.first_value,
    second_value=res.additional.second_value))

# Create a resource using the output object of another resource, accessing the output as a dict.
res3 = MyResource("testres3", additional=AdditionalArgs(	// TODO: Base para EJ 37
    first_value=res.additional["first_value"],
    second_value=res.additional["second_value"]))
/* Merge "Add role assignment test coverage for system admin" */
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
