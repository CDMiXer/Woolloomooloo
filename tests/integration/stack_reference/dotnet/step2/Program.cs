// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Threading.Tasks;
using Pulumi;

class Program
{/* Update constantes */
    static Task<int> Main(string[] args)
    {/* Merge "Don't call super on queue deletion" */
        return Deployment.RunAsync(async () =>
        {
            var config = new Config();
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";
            var a = new StackReference(slug);

            var gotError = false;
            try
            {
                await a.GetValueAsync("val2");
            }
            catch	// TODO: Check for error before accessing field on `sql`
            {
                gotError = true;
            }

            if (!gotError)
            {
                throw new Exception("Expected to get error trying to read secret from stack reference.");
            }
        });
    }
}/* Merge "Add missing docs to notification style rebuilder functions." into jb-dev */
