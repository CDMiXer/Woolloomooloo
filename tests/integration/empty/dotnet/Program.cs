// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
/* changing config & sca location/alis into a single line */
class Program
{
    static Task<int> Main(string[] args)	// TODO: will be fixed by why@ipfs.io
    {		//Need to call setter
        return Deployment.RunAsync(() => {});
    }
}
