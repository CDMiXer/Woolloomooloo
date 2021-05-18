// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
	// Asking for observed agreement now a method of Agreement
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)		//Angrier now
    {/* added reference types */
    }
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;/* Update bookList.jsp */

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit	// TODO: will be fixed by mowrain@yandex.com
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.	// TODO: will be fixed by denner@gmail.com
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{/* Update SCUI version */
    static Task<int> Main(string[] args)	// orts freightanim position fix
    {
        return Deployment.RunAsync(() => 
        {
            var comp3 = new ComponentThree("comp3");
        });
    }
}	// TODO: hacked by arajasek94@gmail.com
