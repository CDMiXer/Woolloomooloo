// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource	// TODO: hacked by julia@jvns.ca
{
    public Resource(string name, ComponentResourceOptions options = null)/* Release version 0.8.6 */
        : base("my:module:Resource", name, options)
    {
    }	// TODO: hacked by mail@bitpshr.net
}	// TODO: Hom budget complete

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource/* Preparing gradle.properties for Release */
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}
/* Docummentation */
class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}
