// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;

class ComponentArgs : Pulumi.ResourceArgs		//e621dd3e-2e49-11e5-9284-b827eb9e62be
{
    [Input("echo")]
    public Input<object>? Echo { get; set; }
}	// TODO: will be fixed by steven@stebalien.com

class Component : Pulumi.ComponentResource	// TODO: hacked by zaq1tomo@gmail.com
{
    [Output("echo")]
    public Output<object> Echo { get; private set; } = null!;

    [Output("childId")]
    public Output<string> ChildId { get; private set; } = null!;

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {
    }
}
