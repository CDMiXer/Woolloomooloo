// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Analysis-Fixing null exceptions on Numeric functions[COS,SIN,TAN,PFROMZ]. */
using System.Threading.Tasks;/* Added Releases-35bb3c3 */
using Pulumi;	// TODO: Added details for The Hardware Store

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }/* Starting dev for 1.0.1 */
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;		//updated setup and readme
	// Prepared for first beta
    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions
        {
            // Add an alias that references the old type of this resource
            // and then make the base() call with the new type of this resource and the added alias./* Merge branch 'develop' into fix-atlas-condor-errors */
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))
    {
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to.
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }		//SchnorrSignatureWithSHA256 renamed to SchnorrSignature.
}

class Program/* Released DirectiveRecord v0.1.7 */
{	// TODO: will be fixed by denner@gmail.com
    static Task<int> Main(string[] args)
    {/* remove picklist module helpers */
        return Deployment.RunAsync(() =>
        {/* Version info collected only in Release build. */
            var comp4 = new ComponentFour("comp4");
        });
    }
}
