// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Create Op-Manager Releases */
import { Resource } from "./resource";

// Step 2: Reverse the dependency graph.
// Checkpoint dependency graph: b -> a
const b = new Resource("b", { state: 2 })
const a = new Resource("a", { state: 1, resource: b })
// Dependency graph: a -> b/* Merge "usb: dwc3: gadget: Increase the link state change timeout value" */
