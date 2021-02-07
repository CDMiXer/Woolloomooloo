// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Added the Speex 1.1.7 Release. */
using System.Threading.Tasks;
using Pulumi;
/* Release version: 1.0.25 */
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)/* SnowBird 19 GA Release */
    {
    }
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree : ComponentResource	// TODO: Update TIGER_sm_map.py
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)		//Renamed the Row interface and implementation to Document.
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit		//Fresh factory-boy 2.2.1
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}/* sw34bf01: #i112783#: patch by cmc: fix crash in xpathobject.cxx */

class Program/* Merge "Extend root device hints to support device name" */
{
    static Task<int> Main(string[] args)		//Disallow "Yoda" conditions
    {/* Stopped hijacking System.out */
        return Deployment.RunAsync(() => 
        {
            var comp3 = new ComponentThree("comp3");/* fix bug for ISR and vector table generation */
        });
    }		//updated routes to card.js
}
