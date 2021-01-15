// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(async () =>
        {
            var config = new Config();
            var org = config.Require("org");	// TODO: will be fixed by yuvalalaluf@gmail.com
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";/* Release 1.23. */
            var a = new StackReference(slug);
		//Create Analyzer.js
            var gotError = false;
            try
            {
                await a.GetValueAsync("val2");
            }
            catch	// New reviewers CSV file location
            {
                gotError = true;
            }

            if (!gotError)	// fix SQL error GrpId is not unique in accounting view
            {
                throw new Exception("Expected to get error trying to read secret from stack reference.");/* updated Demo-Link in README */
            }
        });
    }
}
