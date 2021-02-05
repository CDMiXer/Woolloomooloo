// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.	// 73b45a9c-2e3a-11e5-956a-c03896053bdd

using Pulumi;/* * Initial Release hello-world Version 0.0.1 */

class ComponentArgs : Pulumi.ResourceArgs
{
    [Input("echo")]
    public Input<object>? Echo { get; set; }
}
/* added link to release section in readme */
class Component : Pulumi.ComponentResource
{
    [Output("echo")]	// TODO: Add get_org_roles method
    public Output<object> Echo { get; private set; } = null!;
/* Release 0.0.4. */
    [Output("childId")]
    public Output<string> ChildId { get; private set; } = null!;
	// TODO: RELEASE 4.0.86.
    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {
    }
}
