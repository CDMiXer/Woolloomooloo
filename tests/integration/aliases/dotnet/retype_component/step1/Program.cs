// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
/* Merge "Add metadata for RH Release" */
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {/* Made the SocketService. */
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;		//removes ENV secrets from configration
	// Корректировка в выводе поля Отчество в админке
    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });		//aEqozYxjuSuwLoIylyZZHngOQGoczLli
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}
