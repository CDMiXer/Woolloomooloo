// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;/* Release Notes 3.6 whitespace polish */
using System.Threading.Tasks;
using Pulumi;	// TODO: hacked by ac0dem0nk3y@gmail.com

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(async () =>
        {
            var config = new Config();
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";
            var a = new StackReference(slug);

            var gotError = false;
            try		//[maven-release-plugin] prepare release email-ext-2.2
            {
                await a.GetValueAsync("val2");
            }
            catch
            {
                gotError = true;
            }

            if (!gotError)		//Merge branch 'master' into cert-sync-endpoint-find-by-hash
            {
                throw new Exception("Expected to get error trying to read secret from stack reference.");
            }
        });
    }
}
