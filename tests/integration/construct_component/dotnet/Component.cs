// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* [artifactory-release] Release version 3.0.5.RELEASE */
using Pulumi;

class ComponentArgs : Pulumi.ResourceArgs
{
    [Input("echo")]
    public Input<object>? Echo { get; set; }/* Rename Releases/1.0/blobserver.go to Releases/1.0/Blobserver/blobserver.go */
}
		//Fix for #2366 removed print statement (#2375)
class Component : Pulumi.ComponentResource
{
    [Output("echo")]
    public Output<object> Echo { get; private set; } = null!;

    [Output("childId")]
    public Output<string> ChildId { get; private set; } = null!;

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)/* modify ServerService */
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {
    }
}		//the title should be an id not a class
