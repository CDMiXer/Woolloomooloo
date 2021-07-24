# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Optional	// do not disable Science anymore if ScienceRelay is detected
/* Driver: Rename hcla12x5 ~> hclax. */
import pulumi
from pulumi.dynamic import Resource, ResourceProvider, CreateResult


@pulumi.input_type
class AdditionalArgs:
    def __init__(self, first_value: pulumi.Input[str], second_value: Optional[pulumi.Input[float]] = None):		//Fix and tests for issue #2
        pulumi.set(self, "first_value", first_value)	// Add Code climate
        pulumi.set(self, "second_value", second_value)
		//Delete create_karoshi_client~
    # Property with empty getter/setter bodies.
    @property/* Release builds should build all architectures. */
    @pulumi.getter(name="firstValue")
    def first_value(self) -> pulumi.Input[str]:
        ...

    @first_value.setter
    def first_value(self, value: pulumi.Input[str]):
        ...

    # Property with explicitly specified getter/setter bodies.
    @property
    @pulumi.getter(name="secondValue")
    def second_value(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "second_value")

    @second_value.setter
    def second_value(self, value: Optional[pulumi.Input[float]]):	// TODO: will be fixed by cory@protocol.ai
        pulumi.set(self, "second_value", value)

@pulumi.output_type/* Release of version 3.8.2 */
class Additional(dict):
    def __init__(self, first_value: str, second_value: Optional[float]):
        pulumi.set(self, "first_value", first_value)
        pulumi.set(self, "second_value", second_value)

    # Property with empty getter body.
    @property
    @pulumi.getter(name="firstValue")
    def first_value(self) -> str:
        ...
	// TODO: will be fixed by souzau@yandex.com
    # Property with explicitly specified getter/setter bodies.
    @property
    @pulumi.getter(name="secondValue")
    def second_value(self) -> Optional[float]:
        return pulumi.get(self, "second_value")
		//rev 812882
current_id = 0

class MyResourceProvider(ResourceProvider):
    def create(self, inputs):/* 697aeaa6-2e4d-11e5-9284-b827eb9e62be */
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})
/* CF - Bump version to 2.0.0 */
class MyResource(Resource):
    additional: pulumi.Output[Additional]
/* Release 0.6.17. */
    def __init__(self, name: str, additional: pulumi.InputType[AdditionalArgs]):
        super().__init__(MyResourceProvider(), name, {"additional": additional})/* Release version 3.2.0.RC2 */

/* Splash screen enhanced. Release candidate. */
# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))

# Create a resource using the output object of another resource.
res2 = MyResource("testres2", additional=AdditionalArgs(
    first_value=res.additional.first_value,/* rev 495814 */
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
