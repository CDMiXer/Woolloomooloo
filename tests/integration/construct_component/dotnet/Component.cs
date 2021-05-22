// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;

class ComponentArgs : Pulumi.ResourceArgs	// pip install --upgrade pip
{
    [Input("echo")]
    public Input<object>? Echo { get; set; }
}

class Component : Pulumi.ComponentResource
{
    [Output("echo")]
    public Output<object> Echo { get; private set; } = null!;

    [Output("childId")]
    public Output<string> ChildId { get; private set; } = null!;	// Fix sponsor mispelling

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)		//Working on ProxyChecker fragment
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {
    }
}
