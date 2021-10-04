﻿// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;		//Added additional entity tests

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)	// TODO: Update encdec.php
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component itself.	// TODO: Rename src/Model_ to src/Model/Issue.php
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;/* Release 0.95.030 */

    public ComponentThree(string name, ComponentResourceOptions options = null)/* Reverted Release version */
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name./* [RELEASE] Release version 0.2.0 */
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }/* effective code reviews */
}


class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            // Applying an alias to the instance successfully renames both the component and the children.
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp3" } },
            });
        });
    }
}/* Release 0.0.5 */
