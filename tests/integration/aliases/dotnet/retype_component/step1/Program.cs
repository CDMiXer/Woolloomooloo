// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// service: get options name from string

using System.Threading.Tasks;	// TODO: hacked by timnugent@gmail.com
using Pulumi;

class Resource : ComponentResource
{/* finish test mock */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {	// fix "concern" typo!
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;/* Folder structure of biojava3 project adjusted to requirements of ReleaseManager. */

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
}    
}		//83c38dc0-2e73-11e5-9284-b827eb9e62be

class Program
{
    static Task<int> Main(string[] args)
    {		//updated readme to show new port in runserver.py
        return Deployment.RunAsync(() => 
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}	// TODO: hacked by zaq1tomo@gmail.com
