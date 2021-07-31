// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.	// TODO: will be fixed by yuvalalaluf@gmail.com

using Pulumi;

class ComponentArgs : Pulumi.ResourceArgs
{
    [Input("echo")]
    public Input<object>? Echo { get; set; }/* Update enable-ssh-user-login-other-than-root.md */
}/* Portal Release */
/* Merge "docs: Android for Work updates to DP2 Release Notes" into mnc-mr-docs */
class Component : Pulumi.ComponentResource
{	// TODO: Merge "Removing redundant code in vp9_entropymode.c." into experimental
    [Output("echo")]
    public Output<object> Echo { get; private set; } = null!;

    [Output("childId")]	// Remove unnecessary attribute from example
    public Output<string> ChildId { get; private set; } = null!;

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {
    }
}
