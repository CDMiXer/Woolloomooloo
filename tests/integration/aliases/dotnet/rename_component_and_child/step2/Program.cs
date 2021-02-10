// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;		//Inital Example Project
using Pulumi;
	// [RTM] Drop ListStore stuff
class Resource : ComponentResource/* Unchaining WIP-Release v0.1.41-alpha */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}
		//Translated all [name fields] into Spanish
// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)	// TODO: Fixing Listener priority
    {
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions/* Release of eeacms/www-devel:20.12.5 */
        { 		//added missing new class State
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },
        });
    }	// TODO: #2 updated cids_reference.sql dump script
}
	// Update config.in
class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp5" } },
            });
        });	// TODO: Changed smooth factor to array
    }
}
