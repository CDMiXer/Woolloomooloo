// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* turn on msm_otg for ycable testing */
		//Update sidekiq to version 5.0.1
using System.Threading.Tasks;/* Added Map tests */
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}/* First Release- */

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{		//botlib: uncrustify
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {	// TODO: CMS update of messaging/services/service-number-add by dprothero@twilio.com
        return Deployment.RunAsync(() => 
        {
            var comp4 = new ComponentFour("comp4");
        });/* Fixed bug with the immutable api */
    }/* Merge "Update FSTrigger plugin" */
}
