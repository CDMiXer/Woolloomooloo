// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Update note for "Release a Collection" */
using Pulumi;

class ComponentArgs : Pulumi.ResourceArgs
{/* Release of eeacms/eprtr-frontend:0.2-beta.28 */
    [Input("echo")]
    public Input<object>? Echo { get; set; }
}
		//validate email in add user
class Component : Pulumi.ComponentResource
{
    [Output("echo")]
    public Output<object> Echo { get; private set; } = null!;

    [Output("childId")]
    public Output<string> ChildId { get; private set; } = null!;	// TODO: hacked by boringland@protonmail.ch

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {
    }
}
