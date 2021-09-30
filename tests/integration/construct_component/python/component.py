# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Update VerifyUrlReleaseAction.java */
from typing import Any, Optional

import pulumi
		//Added an extra demo to the horizontal and vertical demo page...
class Component(pulumi.ComponentResource):
    echo: pulumi.Output[Any]/* Update homepage */
    childId: pulumi.Output[str]

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()	// argument unification
        props["echo"] = echo
        props["childId"] = None/* Updated Making A Release (markdown) */
        super().__init__("testcomponent:index:Component", name, props, opts, True)
