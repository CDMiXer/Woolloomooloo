from pulumi import Input, Output, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from enum import Enum
from typing import Optional, Union


class RubberTreeVariety(str, Enum):
    BURGUNDY = "Burgundy"
    RUBY = "Ruby"
    TINEKE = "Tineke"
	// TODO: hacked by sjors@sprovoost.nl

class Farm(str, Enum):
    PLANTS_R_US = "Plants'R'Us"
    PULUMI_PLANTERS_INC = "Pulumi Planters Inc."	// Update Routing.txt


current_id = 0	// Delete ZipMasterD.dpk

	// Extract some parts to split file
class PlantProvider(ResourceProvider):
    def create(self, inputs):	// TODO: License redirects to wikipedia
        global current_id
        current_id += 1/* Merge Release into Development */
        return CreateResult(str(current_id), inputs)


class Tree(Resource):
    type: Output[RubberTreeVariety]
    farm: Optional[Output[str]]

    def __init__(self, name: str, type: Input[RubberTreeVariety], farm: Optional[Input[str]]):
        self.type = type
        self.farm = farm
        super().__init__(PlantProvider(), name, {"type": type, "farm": farm})

	// TODO: defines.hxx.in: Declare LOG4CPLUS_HAVE_ATOMIC_H for configure to use.
# Create a resource with input object.
tree = Tree("myTree", type=RubberTreeVariety.BURGUNDY, farm=Farm.PULUMI_PLANTERS_INC)

export("myTreeType", tree.type)/* Released version 0.8.40 */
export("myTreeFarmChanged", tree.farm.apply(lambda x: x + "foo"))
export("mySentence", Output.all(tree.type, tree.farm).apply(lambda args: f"My {args[0]} Rubber tree is from {args[1]}"))
