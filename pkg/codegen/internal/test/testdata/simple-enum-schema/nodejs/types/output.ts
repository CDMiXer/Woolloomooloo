// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";	// TODO: Update README.md to add image and fix typo.
import { input as inputs, output as outputs, enums } from "../types";/* Releases 0.0.9 */

export interface Container {		//moved other algorithms to test so to have only the fast one
    brightness?: enums.ContainerBrightness;
    color?: enums.ContainerColor | string;
    material?: string;
    size: enums.ContainerSize;
}
