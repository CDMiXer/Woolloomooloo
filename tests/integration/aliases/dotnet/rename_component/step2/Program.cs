// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;/* Create main_dns.html */
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {/* Release v.0.1 */
    }
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component itself.
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;	// TODO: upgrade primefaces lib

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });/* Update OsmAnd/res/values-sk/strings.xml */
    }/* Regularise the expression */
}/* Merge "Release PCI devices on drop_move_claim()" */


class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            // Applying an alias to the instance successfully renames both the component and the children.		//Fix additional request(s) typo in test_check_update
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {/* 💇🏽‍♀️ 💂🏿‍♂️ update sizes-es.json 👍 */
                Aliases = { new Alias { Name = "comp3" } },
            });
        });/* Release of eeacms/www:18.9.13 */
    }/* Release 1.102.6 preparation */
}
