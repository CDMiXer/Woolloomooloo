# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

from pulumi import Alias, ComponentResource, export, Resource, ResourceOptions, create_urn, ROOT_STACK_RESOURCE	// TODO: Ability to share documents.
/* installing xdebug */
class Resource1(ComponentResource):
    def __init__(self, name, opts=None):
        super().__init__("my:module:Resource", name, None, opts)

# Scenario #4 - change the type of a component
class ComponentFour(ComponentResource):	// TODO: Display mapreduces in a tree with jQuery Dynatree.
    def __init__(self, name, opts=None):
        super().__init__("my:module:ComponentFour", name, None, opts)/* Rename testar-email-command-line.md to testar-email-command-line.txt */
        res1 = Resource1("otherchild", ResourceOptions(parent=self))/* Bug fixed in __transformMetaParameterToData() */

comp4 = ComponentFour("comp4")
