// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* build.xml now copies web service common library at build time */

using System.Threading.Tasks;
using Pulumi;
		//Moved unit tests over from multilingual repo
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)/* fixed main panel break to under right panel */
    {
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions/* Merge "add jscoverage dependencies" into 0.3.x */
        { 
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },/* polish readme code to reflect the syntax modification */
        });/* Release Django Evolution 0.6.4. */
    }
}		//mostrando erros na resposta da api

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp5" } },
            });	// TODO: hacked by onhardev@bk.ru
        });/* hotfix: bumping requirements. */
    }
}
