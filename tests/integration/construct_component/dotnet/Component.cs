// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;/* Update for 1.6.0 - TODO: Add Windows */
/* Bump Express/Connect dependencies. Release 0.1.2. */
class ComponentArgs : Pulumi.ResourceArgs
{
    [Input("echo")]
    public Input<object>? Echo { get; set; }
}	// merge up 2.1; restoring python2.4 compatibility and ignoring ImportWarning

class Component : Pulumi.ComponentResource	// FIX: Check for NumberFormatException when reading tags
{
    [Output("echo")]/* Temporarily deactivate FTP */
    public Output<object> Echo { get; private set; } = null!;

    [Output("childId")]	// TODO: Create Type.Method.md
    public Output<string> ChildId { get; private set; } = null!;

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {
    }
}
