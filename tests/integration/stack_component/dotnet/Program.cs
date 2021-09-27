// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack
{	// TODO: will be fixed by ac0dem0nk3y@gmail.com
    [Output("abc")]
    public Output<string> Abc { get; private set; }/* Added option to save a particular file version to disk */

    [Output]
    public Output<int> Foo { get; private set; }
		//Delete bdwn.png
    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }

    public MyStack()
    {		//close MQTT connection on window closed announcement
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
    }
}	// TODO: Merge "Adapting to use the python-saharaclient library"

class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
