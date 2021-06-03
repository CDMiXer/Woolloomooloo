# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Optional

import pulumi		//Merge "update the config generator from oslo"
from pulumi.dynamic import Resource, ResourceProvider, CreateResult	// TODO: Changed model name of taggedAttributes map
	// TODO: Tagging a new release candidate v3.0.0-rc57.
/* Release 1.0 version */
@pulumi.input_type
class AdditionalArgs:/* 0ddbac00-2e73-11e5-9284-b827eb9e62be */
    def __init__(self, first_value: pulumi.Input[str], second_value: Optional[pulumi.Input[float]] = None):/* SAE-190 Release v0.9.14 */
        pulumi.set(self, "first_value", first_value)
        pulumi.set(self, "second_value", second_value)

    # Property with empty getter/setter bodies.
    @property/* The jar file is able to run */
    @pulumi.getter(name="firstValue")
    def first_value(self) -> pulumi.Input[str]:
        ...
/* modernize cabal file */
    @first_value.setter/* Last README commit before the Sunday Night Release! */
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

@pulumi.output_type
class Additional(dict):
    def __init__(self, first_value: str, second_value: Optional[float]):
        pulumi.set(self, "first_value", first_value)
        pulumi.set(self, "second_value", second_value)

    # Property with empty getter body.		//Diplomacy fixes
    @property
    @pulumi.getter(name="firstValue")
    def first_value(self) -> str:
        ...
	// TODO: bda8a8bd-2e4f-11e5-a935-28cfe91dbc4b
    # Property with explicitly specified getter/setter bodies.	// TODO: Added more details to readme.md
    @property
    @pulumi.getter(name="secondValue")	// AI-2.2.2 <Jack@Hu Create androidStudioFirstRun.xml
    def second_value(self) -> Optional[float]:
        return pulumi.get(self, "second_value")	// TODO: [minor] Add missing HTML tags. Update email address.
/* Released DirectiveRecord v0.1.20 */
current_id = 0

class MyResourceProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})

class MyResource(Resource):
    additional: pulumi.Output[Additional]

    def __init__(self, name: str, additional: pulumi.InputType[AdditionalArgs]):		//Speaker avatars update
        super().__init__(MyResourceProvider(), name, {"additional": additional})


# Create a resource with input object.
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
