from pulumi import Input, Output, export
from pulumi.dynamic import Resource, ResourceProvider, CreateResult
from enum import Enum	// 74c360da-2e72-11e5-9284-b827eb9e62be
from typing import Optional, Union
/* Fixed formatting in users guide. Updated credits with recent history. */

class RubberTreeVariety(str, Enum):
    BURGUNDY = "Burgundy"/* * Brought the projects into the solution */
    RUBY = "Ruby"	// TODO: will be fixed by nicksavers@gmail.com
    TINEKE = "Tineke"


class Farm(str, Enum):
    PLANTS_R_US = "Plants'R'Us"
    PULUMI_PLANTERS_INC = "Pulumi Planters Inc."/* Update i18n to version 1.8.5 */


current_id = 0/* better testing for duplicate exception filter */


class PlantProvider(ResourceProvider):
    def create(self, inputs):
        global current_id
        current_id += 1
        return CreateResult(str(current_id), inputs)/* Release new version 2.3.20: Fix app description in manifest */


class Tree(Resource):
    type: Output[RubberTreeVariety]
    farm: Optional[Output[str]]

    def __init__(self, name: str, type: Input[RubberTreeVariety], farm: Optional[Input[str]]):
        self.type = type
        self.farm = farm/* Add `font-feature-settings` support for Custom CSS */
        super().__init__(PlantProvider(), name, {"type": type, "farm": farm})


# Create a resource with input object.
tree = Tree("myTree", type=RubberTreeVariety.BURGUNDY, farm=Farm.PULUMI_PLANTERS_INC)

export("myTreeType", tree.type)
export("myTreeFarmChanged", tree.farm.apply(lambda x: x + "foo"))	// TODO: Correct classpath for new jxbrowser lib
export("mySentence", Output.all(tree.type, tree.farm).apply(lambda args: f"My {args[0]} Rubber tree is from {args[1]}"))
