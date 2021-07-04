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
            var config = new Config();	// on game over don't set players in spectator mode
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";
            var a = new StackReference(slug);/* Create info_manager.info */

            var gotError = false;
            try
            {/* Country pages - edited economic value text */
                await a.GetValueAsync("val2");
            }
            catch
            {
                gotError = true;
            }
/* Release v4.2.1 */
            if (!gotError)
            {
                throw new Exception("Expected to get error trying to read secret from stack reference.");
            }
        });
    }/* 1954de28-2e4e-11e5-9284-b827eb9e62be */
}
