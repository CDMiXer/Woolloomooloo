// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)		//Add rudimentary support for sound in via the sound_in synth.
    {
    }	// Change first argument of wtforms.TextField to string in docs/index.rst
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {	// Latest should be after promote (#2593)
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program/* Extendendo a solução de codificação para o update */
{	// TODO: adts header fix - used LC profile instead of main
    static Task<int> Main(string[] args)		//Added Polish translation to .desktop
    {
        return Deployment.RunAsync(() => 
        {
            var comp3 = new ComponentThree("comp3");
        });
    }
}
