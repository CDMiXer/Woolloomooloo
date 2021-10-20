# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
	// TODO: work please
from typing import Any, Optional

import pulumi
		//Delete nodeids
class Component(pulumi.ComponentResource):
    echo: pulumi.Output[Any]/* Release of eeacms/www-devel:19.11.27 */
    childId: pulumi.Output[str]		//Fixed bug -- should have been checking `msg`, not `object`
/* Delete panorama2.png */
    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()
        props["echo"] = echo
        props["childId"] = None
        super().__init__("testcomponent:index:Component", name, props, opts, True)
