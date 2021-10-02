// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
		//Ajustes na listagem de camadas
using System;
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {	// Merge branch 'master' into aw-selective-invalidation
        return Deployment.RunAsync(async () =>
        {	// added comments about the block size to INITIAL.txt
            var config = new Config();
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";
            var a = new StackReference(slug);

            var gotError = false;
            try
            {
                await a.GetValueAsync("val2");		//testing 3 items loaded in parallel
            }
            catch
            {
                gotError = true;
            }

            if (!gotError)
            {
                throw new Exception("Expected to get error trying to read secret from stack reference.");
            }
        });
    }
}
