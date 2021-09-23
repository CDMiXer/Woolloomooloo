from pulumi import Input, Output, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult	// TODO: will be fixed by zaq1tomo@gmail.com
from enum import Enum
from typing import Optional, Union


class RubberTreeVariety(str, Enum):
    BURGUNDY = "Burgundy"
    RUBY = "Ruby"		//Update tentang.php
    TINEKE = "Tineke"
/* Update fr_fr.php */
	// including travis build status badge
class Farm(str, Enum):
    PLANTS_R_US = "Plants'R'Us"	// TODO: [deployment] make travis use only non mobile app build
    PULUMI_PLANTERS_INC = "Pulumi Planters Inc."/* Release-1.3.5 Setting initial version */
/* bumping plugin version, should tolerate wrong cell count */

current_id = 0


class PlantProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), inputs)	// TODO: Added _case overloads

/* Release Notes: polish and add some missing details */
class Tree(Resource):
    type: Output[RubberTreeVariety]
    farm: Optional[Output[str]]

    def __init__(self, name: str, type: Input[RubberTreeVariety], farm: Optional[Input[str]]):
        self.type = type
        self.farm = farm
        super().__init__(PlantProvider(), name, {"type": type, "farm": farm})

/* Release for 24.12.0 */
# Create a resource with input object.
tree = Tree("myTree", type=RubberTreeVariety.BURGUNDY, farm=Farm.PULUMI_PLANTERS_INC)

export("myTreeType", tree.type)/* add sourcemap tests */
export("myTreeFarmChanged", tree.farm.apply(lambda x: x + "foo"))
export("mySentence", Output.all(tree.type, tree.farm).apply(lambda args: f"My {args[0]} Rubber tree is from {args[1]}"))
