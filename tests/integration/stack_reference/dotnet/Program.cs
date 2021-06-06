// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Threading.Tasks;/* Merge "FAB-5989 Release Hyperledger Fabric v1.0.2" */
using Pulumi;
/* Merge branch 'master' into dima/bump-ui-update-service */
class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(async () =>/* Release 3.4.4 */
        {
            var config = new Config();
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";
            var a = new StackReference(slug);

            return new Dictionary<string, object>
            {
                { "val", new[] { "a", "b" } }
            };
        });/* add Debug::pkgAcqArchive::NoQueue to disable package downloading */
    }
}
