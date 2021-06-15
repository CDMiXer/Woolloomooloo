﻿// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component itself.
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit		//Readme updated with swagger URL
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.	// TODO: Update solving_problems_and_being_lazy.ftl
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });		//popcalendar, adjusted caption for previous month, day, week
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}
	// no duplicate hash entries in settings.yml

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            // Applying an alias to the instance successfully renames both the component and the children.
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp3" } },	// various cleanups of interface to server and and also added scrolling.
            });
        });
    }/* Product-Backlog-475: Move the field in stock */
}
