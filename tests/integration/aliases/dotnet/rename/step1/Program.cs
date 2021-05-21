// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource		//BUG : java.sql.date extends java.util.date.
{	// Create WINNF_FT_S_FPR_testcase.py
    public Resource(string name, ComponentResourceOptions options = null)	// TODO: will be fixed by alex.gaynor@gmail.com
        : base("my:module:Resource", name, options)
{    
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {		//Sep Update
        return Deployment.RunAsync(() => 
        {
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");	// UsuarioServicio
        });
    }
}
