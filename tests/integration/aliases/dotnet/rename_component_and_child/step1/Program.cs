// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
	// TODO: typo minifies => minifiers
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)		//Updated MOD_VERSION in resources
        : base("my:module:Resource", name, options)
    {
    }/* 0.4 Release */
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource		//- maj de vue de la form mot de passe publié
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)	// TODO: Restored proper Sonos Control source that was inadvertently overwritten
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }		//patch .gitignore and .npmignore files
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp5 = new ComponentFive("comp5");
        });
    }		//Create ApplicationUser.cs
}
