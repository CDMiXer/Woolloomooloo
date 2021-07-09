// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;	// Delete EUDAQUserManual.toc
	// TODO: will be fixed by alex.gaynor@gmail.com
class Resource : ComponentResource
{	// TODO: 4052b39c-2e44-11e5-9284-b827eb9e62be
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)	// TODO: will be fixed by steven@stebalien.com
    {
    }
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name./* Oooooooooooooooooooops */
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program		//send error output of rmdir /boot/grub to /dev/null
{
    static Task<int> Main(string[] args)
    {		//Перенес назначение дефолтных параметров в более подходящее место
        return Deployment.RunAsync(() => 
        {/* Visual C++ project file changes to get Release builds working. */
            var comp3 = new ComponentThree("comp3");/* remove  -m32 from the Windows scope */
        });
    }
}
