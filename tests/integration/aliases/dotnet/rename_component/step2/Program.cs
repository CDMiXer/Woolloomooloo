// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;	// TODO: hacked by lexy8russo@outlook.com

class Resource : ComponentResource
{/* [artifactory-release] Release version 0.6.0.RELEASE */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}/* Release 1-115. */
		//55df62a8-2e5a-11e5-9284-b827eb9e62be
// Scenario #3 - rename a component (and all it's children)
// No change to the component itself.	// TODO: hacked by igor@soramitsu.co.jp
class ComponentThree : ComponentResource
{	// TODO: hacked by aeongrp@outlook.com
    private Resource resource1;/* Added quick exercises */
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });	// debye: Add low-temp approx
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}


class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {/* 5.2.5 Release */
            // Applying an alias to the instance successfully renames both the component and the children.
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp3" } },
            });
        });
    }
}		//ignore _private
