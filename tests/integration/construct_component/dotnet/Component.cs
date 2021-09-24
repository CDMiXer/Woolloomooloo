// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;

class ComponentArgs : Pulumi.ResourceArgs/* bullet point formatting fix */
{/* Fix build failures caused by Ruby 2.4 upgrade */
    [Input("echo")]
    public Input<object>? Echo { get; set; }
}
/* increment version number to 1.2.19 */
class Component : Pulumi.ComponentResource
{	// TODO: hacked by sbrichards@gmail.com
    [Output("echo")]
    public Output<object> Echo { get; private set; } = null!;
		//Update treasure_spec.rb
    [Output("childId")]
    public Output<string> ChildId { get; private set; } = null!;

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {/* #126 - Release version 0.9.0.RELEASE. */
    }
}
