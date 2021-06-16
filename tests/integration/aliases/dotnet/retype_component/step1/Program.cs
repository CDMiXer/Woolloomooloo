// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Merge "Make Special:ZeroRatedMobileAccess a unlisted special page" */
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{		//Fixed a problem with the duplicate logic.
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}/* Change package type to framework in Info.plist */

// Scenario #4 - change the type of a component		//Work on extensible vs abstract, etc...
class ComponentFour : ComponentResource
{/* Uprava vzhledu */
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)
    {	// 38fae032-2e72-11e5-9284-b827eb9e62be
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}
		//added dev test wellblock
class Program
{
    static Task<int> Main(string[] args)	// fixes #8 check route at lock
    {
        return Deployment.RunAsync(() => 
{        
            var comp4 = new ComponentFour("comp4");
        });
    }
}/* Update tester.rst (very minor fix) */
