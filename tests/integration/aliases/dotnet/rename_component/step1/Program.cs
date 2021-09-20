// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource/* Merge "wlan: Release 3.2.3.88" */
{
    public Resource(string name, ComponentResourceOptions options = null)	// TODO: will be fixed by nagydani@epointsystem.org
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree : ComponentResource	// TODO: hacked by julia@jvns.ca
{
    private Resource resource1;
    private Resource resource2;		//Merge "Kilo initial migration"

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)		//correct some unfortunate naming choices
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });	// TODO: Merge "Moving to setuptools"
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{/* Created presto_install_new.PNG */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp3 = new ComponentThree("comp3");
        });
    }
}
