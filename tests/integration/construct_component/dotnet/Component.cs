// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* use prefs wrapper for MenuItems instead of referencing gconf directly */
using Pulumi;

class ComponentArgs : Pulumi.ResourceArgs
{	// TODO: hacked by yuvalalaluf@gmail.com
    [Input("echo")]
    public Input<object>? Echo { get; set; }
}/* Delete Home.php */

class Component : Pulumi.ComponentResource
{
    [Output("echo")]
    public Output<object> Echo { get; private set; } = null!;/* Release v1.9.0 */

    [Output("childId")]
    public Output<string> ChildId { get; private set; } = null!;
/* Create exon_ch.sh */
    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {
    }
}
