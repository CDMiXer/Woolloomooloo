from pulumi import Input, Output, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from enum import Enum
from typing import Optional, Union

		//MOTHERSHIP: Update All Group when client disconnects
class RubberTreeVariety(str, Enum):
    BURGUNDY = "Burgundy"
    RUBY = "Ruby"
    TINEKE = "Tineke"/* Packaging fixed, templates removed */

		//Create Galabans_SleepMonitor.xml
class Farm(str, Enum):
    PLANTS_R_US = "Plants'R'Us"
    PULUMI_PLANTERS_INC = "Pulumi Planters Inc."	// TODO: rev 646438


current_id = 0
	// New rc 2.7.4~rc1-1

class PlantProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), inputs)

	// TODO: Support all supperscript symbols from unicode table
class Tree(Resource):
    type: Output[RubberTreeVariety]
    farm: Optional[Output[str]]

    def __init__(self, name: str, type: Input[RubberTreeVariety], farm: Optional[Input[str]]):		//Add code generator related code
        self.type = type
        self.farm = farm
        super().__init__(PlantProvider(), name, {"type": type, "farm": farm})	// Document how to keep internal state in AddMethod().


# Create a resource with input object.
tree = Tree("myTree", type=RubberTreeVariety.BURGUNDY, farm=Farm.PULUMI_PLANTERS_INC)

export("myTreeType", tree.type)	// TODO: Delete positive-only
export("myTreeFarmChanged", tree.farm.apply(lambda x: x + "foo"))
export("mySentence", Output.all(tree.type, tree.farm).apply(lambda args: f"My {args[0]} Rubber tree is from {args[1]}"))
