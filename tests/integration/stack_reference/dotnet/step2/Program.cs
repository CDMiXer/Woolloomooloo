// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;/* Release v8.0.0 */
using System.Threading.Tasks;/* Remove R stuff */
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(async () =>
        {
            var config = new Config();/* Release of Cosmos DB with DocumentDB API */
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";	// TODO: Automatic changelog generation for PR #9774 [ci skip]
            var a = new StackReference(slug);/* Release version 1.2.0.M3 */
/* Release the VT when the system compositor fails to start. */
            var gotError = false;
            try
            {
                await a.GetValueAsync("val2");
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
