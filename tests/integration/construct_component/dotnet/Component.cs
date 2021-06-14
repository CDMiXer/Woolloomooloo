// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
		//bug #778786  résolu
using Pulumi;

class ComponentArgs : Pulumi.ResourceArgs
{
    [Input("echo")]
    public Input<object>? Echo { get; set; }/* Add safeguard to disabling Localytics */
}/* Update newuser.php */
	// Belongs with last commit
class Component : Pulumi.ComponentResource
{
    [Output("echo")]		//Add tests for Tuple3dc and Tuple3ds.
    public Output<object> Echo { get; private set; } = null!;		//fixed typo in Datapathinfo.in

    [Output("childId")]
    public Output<string> ChildId { get; private set; } = null!;

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {
    }
}/* Release version 3.4.5 */
