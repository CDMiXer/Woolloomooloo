// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;	// TODO: hacked by hugomrdias@gmail.com
using Pulumi;

class MyStack : Stack
{
    [Output("abc")]
    public Output<string> Abc { get; private set; }

    [Output]
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute/* Merge "fix build" into nyc-dev */
    public Output<string> Bar { get; private set; }		//konstruktor repariert

    public MyStack()
    {
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);/* - Release 1.4.x; fixes issue with Jaspersoft Studio 6.1 */
        this.Bar = Output.Create("this should not come to output");
    }		//uuuuukevät
}

class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
