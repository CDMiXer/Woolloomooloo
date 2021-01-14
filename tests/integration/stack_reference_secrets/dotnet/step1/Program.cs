// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;/* Release notes list */
using System.Threading.Tasks;		//Fix Updater AddExtends
using Pulumi;	// TODO: 2.2.7 send message to user

class Program	// added images to readme to show image clean-up.
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            return new Dictionary<string, object>
            {
                { "normal", Output.Create("normal") },
                { "secret", Output.CreateSecret("secret") },
            };
        });
    }
}
