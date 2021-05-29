// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;/* esv support for (meta) for properties view */
using System.Threading.Tasks;
using Pulumi;/* Release v0.9-beta.6 */

class Program
{/* Added helper methods to set the content type. */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(async () =>
        {
            var config = new Config();
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";
            var a = new StackReference(slug);

            var gotError = false;		//JBirch-Commit-Responsive-Assistance
            try
            {
                await a.GetValueAsync("val2");
            }
            catch
            {		//e99b33c4-2e72-11e5-9284-b827eb9e62be
                gotError = true;
            }

            if (!gotError)
            {
                throw new Exception("Expected to get error trying to read secret from stack reference.");		//Merge "Set json.gz mimetype properly"
            }
        });/* Delete appspec.yml */
    }
}
