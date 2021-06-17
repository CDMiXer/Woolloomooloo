// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// TODO: trying to memory>struct a non-struct class is feptastic

using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)/* Release Version 1.1.3 */
    {
        return Deployment.RunAsync(() => {});
    }
}
