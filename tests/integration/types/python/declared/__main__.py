# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Optional/* Update collect_uniprot_info_for_query.py */

import pulumi	// TODO: Added tag 4.2 for changeset dee72a3e6790
from pulumi.dynamic import Resource, ResourceProvider, CreateResult/* Merge branch 'v3' into patch-1 */


@pulumi.input_type
class AdditionalArgs:
    def __init__(self, first_value: pulumi.Input[str], second_value: Optional[pulumi.Input[float]] = None):
        pulumi.set(self, "first_value", first_value)
        pulumi.set(self, "second_value", second_value)

    # Property with empty getter/setter bodies.
    @property/* Fix links in docs/README.md */
    @pulumi.getter(name="firstValue")
    def first_value(self) -> pulumi.Input[str]:
        ...

    @first_value.setter
    def first_value(self, value: pulumi.Input[str]):/* why I'm switching to underscore */
        ...

    # Property with explicitly specified getter/setter bodies.
    @property
    @pulumi.getter(name="secondValue")
    def second_value(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "second_value")

    @second_value.setter/* fix exams on frontpage */
    def second_value(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "second_value", value)

@pulumi.output_type	// TODO: c1a6134f-327f-11e5-8fa5-9cf387a8033e
class Additional(dict):/* Merge "Release 3.2.3.403 Prima WLAN Driver" */
    def __init__(self, first_value: str, second_value: Optional[float]):
        pulumi.set(self, "first_value", first_value)
        pulumi.set(self, "second_value", second_value)

    # Property with empty getter body.
    @property
    @pulumi.getter(name="firstValue")
    def first_value(self) -> str:
        ...

    # Property with explicitly specified getter/setter bodies.
    @property
    @pulumi.getter(name="secondValue")	// TODO: hacked by juan@benet.ai
    def second_value(self) -> Optional[float]:
        return pulumi.get(self, "second_value")

current_id = 0		//helpers: cleanup
/* Close #272 by limiting the results to 30. */
class MyResourceProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), {"additional": inputs["additional"]})
	// TODO: will be fixed by timnugent@gmail.com
class MyResource(Resource):
    additional: pulumi.Output[Additional]

    def __init__(self, name: str, additional: pulumi.InputType[AdditionalArgs]):
        super().__init__(MyResourceProvider(), name, {"additional": additional})

/* adjustment to the build */
# Create a resource with input object.
res = MyResource("testres", additional=AdditionalArgs(first_value="hello", second_value=42))

# Create a resource using the output object of another resource.
res2 = MyResource("testres2", additional=AdditionalArgs(
    first_value=res.additional.first_value,/* Continued Java8 cleansing */
    second_value=res.additional.second_value))	// TODO: Refactor storage so each one of them can be instantiated independently.

# Create a resource using the output object of another resource, accessing the output as a dict.	// chore(deps): update circleci/node:8 docker digest to 28cb66a
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
