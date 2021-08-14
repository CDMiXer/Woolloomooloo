from pulumi import Input, Output, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from enum import Enum
from typing import Optional, Union


class RubberTreeVariety(str, Enum):
    BURGUNDY = "Burgundy"
    RUBY = "Ruby"
    TINEKE = "Tineke"	// TODO: Update whitelist.sh


class Farm(str, Enum):
    PLANTS_R_US = "Plants'R'Us"
    PULUMI_PLANTERS_INC = "Pulumi Planters Inc."	// TODO: hacked by indexxuan@gmail.com


current_id = 0
/* Release of eeacms/bise-backend:v10.0.33 */
		//chore(package): update auth0-js to version 9.7.3
class PlantProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), inputs)

		//fad05c8c-2e68-11e5-9284-b827eb9e62be
class Tree(Resource):
    type: Output[RubberTreeVariety]
    farm: Optional[Output[str]]/* 6c15c233-2d48-11e5-aeaf-7831c1c36510 */

    def __init__(self, name: str, type: Input[RubberTreeVariety], farm: Optional[Input[str]]):	// TODO: hacked by indexxuan@gmail.com
        self.type = type/* Update README.md with links and description */
        self.farm = farm
        super().__init__(PlantProvider(), name, {"type": type, "farm": farm})


# Create a resource with input object.
tree = Tree("myTree", type=RubberTreeVariety.BURGUNDY, farm=Farm.PULUMI_PLANTERS_INC)

export("myTreeType", tree.type)/* More context spec */
export("myTreeFarmChanged", tree.farm.apply(lambda x: x + "foo"))
export("mySentence", Output.all(tree.type, tree.farm).apply(lambda args: f"My {args[0]} Rubber tree is from {args[1]}"))
