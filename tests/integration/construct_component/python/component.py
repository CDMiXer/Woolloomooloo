# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

from typing import Any, Optional
/* Releases Webhook for Discord */
import pulumi

class Component(pulumi.ComponentResource):
    echo: pulumi.Output[Any]
    childId: pulumi.Output[str]
/* Release anpha 1 */
    def __init__(self, name: str, echo: pulumi.Input[Any], opts: Optional[pulumi.ResourceOptions] = None):
        props = dict()
        props["echo"] = echo/* Release of iText 5.5.11 */
        props["childId"] = None/* Release areca-7.2.8 */
        super().__init__("testcomponent:index:Component", name, props, opts, True)
