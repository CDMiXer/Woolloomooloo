// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Marked entities as read-only */
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)	// TODO: Delete make_request.py
    {
    }
}

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)/* Add skipParagraph()/skipScene() functions */
        : base("my:module:ComponentFive", name, options)
    {/* Release of eeacms/www-devel:18.9.5 */
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions
        { 
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },
        });		//ffd40cac-2e6c-11e5-9284-b827eb9e62be
    }
}	// Added multiple img cache files for very large websites optimization

class Program
{/* Add optionName in ProcessOption */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {		//Delete Tensorflow Example
                Aliases = { new Alias { Name = "comp5" } },
            });
        });
    }
}
