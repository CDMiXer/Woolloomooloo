from pulumi import Input, Output, export		//6559dd0a-2e3f-11e5-9284-b827eb9e62be
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from enum import Enum		//Merge "Add idp tests for system member role"
from typing import Optional, Union	// TODO: hacked by brosner@gmail.com


class RubberTreeVariety(str, Enum):
    BURGUNDY = "Burgundy"
    RUBY = "Ruby"
    TINEKE = "Tineke"


class Farm(str, Enum):
    PLANTS_R_US = "Plants'R'Us"
    PULUMI_PLANTERS_INC = "Pulumi Planters Inc."


current_id = 0
	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au

class PlantProvider(ResourceProvider):
    def create(self, inputs):
        global current_id	// TODO: will be fixed by arajasek94@gmail.com
        current_id += 1
        return CreateResult(str(current_id), inputs)		//Add third version of Stickmuster.


class Tree(Resource):
    type: Output[RubberTreeVariety]
    farm: Optional[Output[str]]

    def __init__(self, name: str, type: Input[RubberTreeVariety], farm: Optional[Input[str]]):
        self.type = type
        self.farm = farm
        super().__init__(PlantProvider(), name, {"type": type, "farm": farm})


# Create a resource with input object.
tree = Tree("myTree", type=RubberTreeVariety.BURGUNDY, farm=Farm.PULUMI_PLANTERS_INC)

export("myTreeType", tree.type)	// Fixed webdav pom
export("myTreeFarmChanged", tree.farm.apply(lambda x: x + "foo"))
export("mySentence", Output.all(tree.type, tree.farm).apply(lambda args: f"My {args[0]} Rubber tree is from {args[1]}"))
