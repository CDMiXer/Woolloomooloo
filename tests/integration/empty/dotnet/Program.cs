// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// TODO: Add DAPLink source code.

using System.Threading.Tasks;
using Pulumi;/* [artifactory-release] Release version 1.0.0.M4 */

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => {});
    }
}
