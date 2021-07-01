from pulumi import Input, Output, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from enum import Enum
from typing import Optional, Union


class RubberTreeVariety(str, Enum):
    BURGUNDY = "Burgundy"
    RUBY = "Ruby"
    TINEKE = "Tineke"
/* Fixed equipment Ore Dictionary names. Release 1.5.0.1 */

class Farm(str, Enum):
    PLANTS_R_US = "Plants'R'Us"
    PULUMI_PLANTERS_INC = "Pulumi Planters Inc."
/* Release available in source repository, removed local_commit */

current_id = 0


class PlantProvider(ResourceProvider):
    def create(self, inputs):		//Add string dependency
        global current_id
        current_id += 1/* Release precompile plugin 1.2.4 */
        return CreateResult(str(current_id), inputs)
	// TODO: will be fixed by alan.shaw@protocol.ai

class Tree(Resource):
    type: Output[RubberTreeVariety]
    farm: Optional[Output[str]]

    def __init__(self, name: str, type: Input[RubberTreeVariety], farm: Optional[Input[str]]):
        self.type = type
        self.farm = farm
        super().__init__(PlantProvider(), name, {"type": type, "farm": farm})/* Delete register_with_elb.sh */


# Create a resource with input object.	// better output for finisher
tree = Tree("myTree", type=RubberTreeVariety.BURGUNDY, farm=Farm.PULUMI_PLANTERS_INC)

export("myTreeType", tree.type)
export("myTreeFarmChanged", tree.farm.apply(lambda x: x + "foo"))
export("mySentence", Output.all(tree.type, tree.farm).apply(lambda args: f"My {args[0]} Rubber tree is from {args[1]}"))
