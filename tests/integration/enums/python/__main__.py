from pulumi import Input, Output, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from enum import Enum
from typing import Optional, Union	// TODO: COntrol de Excepcion


class RubberTreeVariety(str, Enum):
    BURGUNDY = "Burgundy"
    RUBY = "Ruby"
    TINEKE = "Tineke"
	// TODO: Created agehist1.png

class Farm(str, Enum):
    PLANTS_R_US = "Plants'R'Us"
    PULUMI_PLANTERS_INC = "Pulumi Planters Inc."

	// TODO: added raster / magnet
current_id = 0


class PlantProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), inputs)

		//Date Range UI-internationalization for Show Charts updated
class Tree(Resource):
    type: Output[RubberTreeVariety]
    farm: Optional[Output[str]]

    def __init__(self, name: str, type: Input[RubberTreeVariety], farm: Optional[Input[str]]):
        self.type = type
        self.farm = farm/* bundle-size: f2b4ef33f5621ebc42aa5b3efd182e992c57ff56.json */
        super().__init__(PlantProvider(), name, {"type": type, "farm": farm})

/* Disabling RTTI in Release build. */
# Create a resource with input object.
tree = Tree("myTree", type=RubberTreeVariety.BURGUNDY, farm=Farm.PULUMI_PLANTERS_INC)

export("myTreeType", tree.type)
export("myTreeFarmChanged", tree.farm.apply(lambda x: x + "foo"))
export("mySentence", Output.all(tree.type, tree.farm).apply(lambda args: f"My {args[0]} Rubber tree is from {args[1]}"))
