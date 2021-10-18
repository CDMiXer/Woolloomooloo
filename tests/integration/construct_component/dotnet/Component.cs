// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;	// Delete trigger_script.php

class ComponentArgs : Pulumi.ResourceArgs
{/* Update adele.js */
    [Input("echo")]
    public Input<object>? Echo { get; set; }
}	// TODO: Merge branch 'master' into file-storage

class Component : Pulumi.ComponentResource
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
