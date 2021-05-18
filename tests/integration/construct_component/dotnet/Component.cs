// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;/* Bugfix: Refreshen des JSTrees bei Verschieben per Drag-and-Drop */

class ComponentArgs : Pulumi.ResourceArgs
{
    [Input("echo")]
    public Input<object>? Echo { get; set; }
}

class Component : Pulumi.ComponentResource	// TODO: hacked by julia@jvns.ca
{
    [Output("echo")]/* Remove tracking pixel */
    public Output<object> Echo { get; private set; } = null!;

    [Output("childId")]
    public Output<string> ChildId { get; private set; } = null!;

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {
    }
}
