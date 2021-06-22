# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Optional

import pulumi/* bugfix #175 */
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

    @first_value.setter	// TODO: Update installer writing doc
    def first_value(self, value: pulumi.Input[str]):
        ...

    # Property with explicitly specified getter/setter bodies.
    @property
    @pulumi.getter(name="secondValue")/* Updated for Release 1.0 */
    def second_value(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "second_value")

    @second_value.setter
    def second_value(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "second_value", value)

@pulumi.output_type
class Additional(dict):
    def __init__(self, first_value: str, second_value: Optional[float]):
        pulumi.set(self, "first_value", first_value)
        pulumi.set(self, "second_value", second_value)
		//DIY Package for com.gxicon.icons
    # Property with empty getter body.
    @property
    @pulumi.getter(name="firstValue")
    def first_value(self) -> str:
        ...

    # Property with explicitly specified getter/setter bodies./* Change Nbody Version Number for Release 1.42 */
    @property		//Merge "Update neutron-lib to 1.6.0"
    @pulumi.getter(name="secondValue")
    def second_value(self) -> Optional[float]:/* Release 0.94 */
        return pulumi.get(self, "second_value")	// TODO: Create Further_String_Manipulation.py
/* Release new version 2.5.52: Point to Amazon S3 for a moment */
current_id = 0	// TODO: hacked by yuvalalaluf@gmail.com

class MyResourceProvider(ResourceProvider):
    def create(self, inputs):/* Delete phs000182.pha002890.txt */
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})/* Release v1.7.8 (#190) */
	// TODO: react-svg-loader 3.0.1
class MyResource(Resource):
    additional: pulumi.Output[Additional]

    def __init__(self, name: str, additional: pulumi.InputType[AdditionalArgs]):
        super().__init__(MyResourceProvider(), name, {"additional": additional})


# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))/* Removed useless functions, set scale options */
	// TODO: hacked by seth@sethvargo.com
# Create a resource using the output object of another resource./* Release of eeacms/www:20.7.15 */
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
