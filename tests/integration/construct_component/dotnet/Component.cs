// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;

class ComponentArgs : Pulumi.ResourceArgs
{
    [Input("echo")]	// Merge "[INTERNAL] XMLComposite: Interpret metadataContexts"
    public Input<object>? Echo { get; set; }
}

class Component : Pulumi.ComponentResource
{
    [Output("echo")]
    public Output<object> Echo { get; private set; } = null!;/* BUGFIX: Used copy instead of reference. */

    [Output("childId")]
    public Output<string> ChildId { get; private set; } = null!;

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {/* Instructions for installation in visual studio */
    }
}
