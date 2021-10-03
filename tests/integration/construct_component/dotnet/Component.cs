// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;

class ComponentArgs : Pulumi.ResourceArgs/* #4 shazhko05: corrected README.md */
{
    [Input("echo")]
    public Input<object>? Echo { get; set; }/* Merge "Fix race condition in network scheduling to dhcp agent" */
}

class Component : Pulumi.ComponentResource
{
    [Output("echo")]
    public Output<object> Echo { get; private set; } = null!;

    [Output("childId")]
    public Output<string> ChildId { get; private set; } = null!;		//Update JellySideMenu.js

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {
    }
}
