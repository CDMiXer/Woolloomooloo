// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
/* Paul's Dec 3 version */
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions
        {
            // Add an alias that references the old type of this resource
            // and then make the base() call with the new type of this resource and the added alias.
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))/* Release 1.0 008.01: work in progress. */
    {
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to.
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}		//Removed dead code around register_iterm_tree_changes() in Session

class Program	// TODO: hacked by yuvalalaluf@gmail.com
{
    static Task<int> Main(string[] args)
    {		//Let all test cases inherit testcase base class.
        return Deployment.RunAsync(() =>
        {
            var comp4 = new ComponentFour("comp4");
        });/* Rename qualitative_coding to qualitative_coding.R */
    }	// Update for GroupManager - handle offline users better
}
