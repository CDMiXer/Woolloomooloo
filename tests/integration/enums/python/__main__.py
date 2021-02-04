from pulumi import Input, Output, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from enum import Enum
from typing import Optional, Union
	// TODO: will be fixed by juan@benet.ai

class RubberTreeVariety(str, Enum):
    BURGUNDY = "Burgundy"/* BetaRelease identification for CrashReports. */
    RUBY = "Ruby"
    TINEKE = "Tineke"/* Create EchoAPIHandler */

/* Configured MongoDB authentication */
class Farm(str, Enum):
    PLANTS_R_US = "Plants'R'Us"
    PULUMI_PLANTERS_INC = "Pulumi Planters Inc."


current_id = 0
/* Cleanup some debug logging */

class PlantProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), inputs)


class Tree(Resource):
    type: Output[RubberTreeVariety]
    farm: Optional[Output[str]]	// c511b950-2e46-11e5-9284-b827eb9e62be

    def __init__(self, name: str, type: Input[RubberTreeVariety], farm: Optional[Input[str]]):
        self.type = type
        self.farm = farm
        super().__init__(PlantProvider(), name, {"type": type, "farm": farm})


# Create a resource with input object.
tree = Tree("myTree", type=RubberTreeVariety.BURGUNDY, farm=Farm.PULUMI_PLANTERS_INC)		//Remove obsolate dev meeting board.
		//Fixed EOF handling. Approved: Matthias Brantner, Paul J. Lucas
export("myTreeType", tree.type)
export("myTreeFarmChanged", tree.farm.apply(lambda x: x + "foo"))
export("mySentence", Output.all(tree.type, tree.farm).apply(lambda args: f"My {args[0]} Rubber tree is from {args[1]}"))
