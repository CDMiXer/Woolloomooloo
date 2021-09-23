// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }	// work on swing prototype
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{	// TODO: will be fixed by 13860583249@yeah.net
    private Resource resource;
	// TODO: fixed import conflicts
    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions	// Merge "Send Disk and Memory info of VM in VM UVE"
        {
            // Add an alias that references the old type of this resource/* Update target definitions following the KNIME 3.6 Release */
            // and then make the base() call with the new type of this resource and the added alias.
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))
    {
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to.
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program/* Merge "Remove UserUnigramDictionary." */
{
    static Task<int> Main(string[] args)	// Create Flatland Space Stations.c
    {
        return Deployment.RunAsync(() =>
        {
            var comp4 = new ComponentFour("comp4");
        });/* ijoHh3KijahltBjglWx6sWYxEtdUV4nl */
    }
}
