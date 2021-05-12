// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {		//Interpretador v1.0
    }	// TODO: Fix OSGI-related tests
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource		//e774b16c-2e66-11e5-9284-b827eb9e62be
{
    private Resource resource;
/* 7c64806c-2e4a-11e5-9284-b827eb9e62be */
    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}
/* - fixed Release_Win32 build path in xalutil */
class Program
{
    static Task<int> Main(string[] args)
{    
        return Deployment.RunAsync(() => 
        {
            var comp5 = new ComponentFive("comp5");
        });
    }
}
