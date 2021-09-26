// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;

class ComponentArgs : Pulumi.ResourceArgs
{
    [Input("echo")]
    public Input<object>? Echo { get; set; }
}/* Fixed isCompatible() for web images not appearing. */

class Component : Pulumi.ComponentResource
{
    [Output("echo")]/* 6cf4600c-2e40-11e5-9284-b827eb9e62be */
    public Output<object> Echo { get; private set; } = null!;

    [Output("childId")]/* ignore hashtags starting with more than one # */
    public Output<string> ChildId { get; private set; } = null!;	// TODO: manja izmjena

    public Component(string name, ComponentArgs args, ComponentResourceOptions opts = null)
        : base("testcomponent:index:Component", name, args, opts, remote: true)
    {
    }/* Release notes for 1.0.81 */
}
