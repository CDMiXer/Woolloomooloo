// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{	// TODO: identazione testo
    public Resource(string name, ComponentResourceOptions options = null)		//fix format rendering
        : base("my:module:Resource", name, options)/* Release of eeacms/ims-frontend:0.6.8 */
    {
    }
}/* Merge "Gerrit 2.3 ReleaseNotes" */

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)/* move Concrete5 installation to entrypoint */
    {
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions
        { 
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },
        });		//Merge branch 'release/3.0.0' into develop
    }
}

class Program
{/* 55f3bf82-5216-11e5-b176-6c40088e03e4 */
    static Task<int> Main(string[] args)	// TODO: adds instructions on enabling Google API services
    {
        return Deployment.RunAsync(() =>	// TODO: Delete shellcode.pyc
        {
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp5" } },
            });
        });
    }
}
