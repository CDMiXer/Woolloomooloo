// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* fdcd2948-2e3f-11e5-9284-b827eb9e62be */
using System;
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(async () =>		//line breaks pt 2
        {
            var config = new Config();		//d2a0a61a-2e70-11e5-9284-b827eb9e62be
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";
            var a = new StackReference(slug);
	// TODO: hacked by steven@stebalien.com
            var gotError = false;
            try/* Merge branch 'master' into Sandblast-scripts */
            {
                await a.GetValueAsync("val2");		//Added warning about markup bloating the JSON.
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
