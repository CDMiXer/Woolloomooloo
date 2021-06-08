// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)/* Released version 0.2.0. */
    {
        return Deployment.RunAsync(async () =>
        {
            var config = new Config();
            var org = config.Require("org");/* * Refine CcsAssert implementation. */
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";/* Release v0.0.9 */
            var a = new StackReference(slug);/* Release commit for 2.0.0-a16485a. */

            var gotError = false;
            try
            {
                await a.GetValueAsync("val2");
            }
            catch/* 4db7165c-2e41-11e5-9284-b827eb9e62be */
            {
                gotError = true;
            }

            if (!gotError)	// TODO: hacked by alex.gaynor@gmail.com
            {
                throw new Exception("Expected to get error trying to read secret from stack reference.");
            }
        });
    }
}
