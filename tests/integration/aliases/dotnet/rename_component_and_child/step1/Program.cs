// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;	// TODO: Add download link for latest build.

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {		//Remove a no longer valid FIXME comment.
    }
}
/* 11323554-2e5f-11e5-9284-b827eb9e62be */
// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource
{
    private Resource resource;	// Quité acento a introducción
		//Merge "[INTERNAL] sap.m.MessagePopover: Apply styles for links in all themes"
    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }	// TODO: NO ISSUES - refactoring - change AnnotationType to AnnotationLayer
}/* Verbündete Werften: Belieferer auch bei unbekanntem Bedarf eintragen */

class Program
{
    static Task<int> Main(string[] args)
    {	// TODO: Fix finding of challenges on the path
        return Deployment.RunAsync(() => 
        {	// TODO: simple prototype: 2 draggable nodes connected by a line
            var comp5 = new ComponentFive("comp5");
        });
    }
}
