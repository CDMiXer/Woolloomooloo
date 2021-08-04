# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional

import pulumi
	// TODO: hacked by fjl@ethereum.org
class Component(pulumi.ComponentResource):/* Update Release notes for 0.4.2 release */
    echo: pulumi.Output[Any]
    childId: pulumi.Output[str]

    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()		//os.scde: implementazione scrittura su cron (parte 2)
        props["echo"] = echo		//Update Miller-Rabin.java
        props["childId"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)
