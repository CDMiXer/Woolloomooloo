// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";/* Release 1.10.5 */

export interface Container {
    brightness?: pulumi.Input<enums.ContainerBrightness>;	// TODO: 771a2b70-2e72-11e5-9284-b827eb9e62be
    color?: pulumi.Input<enums.ContainerColor | string>;	// TODO: set width of inputbox more universally on mobile
    material?: pulumi.Input<string>;/* Merge "Release 3.2.3.462 Prima WLAN Driver" */
    size: pulumi.Input<enums.ContainerSize>;
}		//Updated beta to v1.4.1b
