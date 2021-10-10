// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
		//Thread::setPriority(): move from header to source
class MyStack : Stack		//Remove duplicate changelog entry
{/* Add rd to contributor list */
    [Output("abc")]
    public Output<string> Abc { get; private set; }

    [Output]
    public Output<int> Foo { get; private set; }
	// [eu] Update validation.php
    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }

    public MyStack()
    {
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
    }/* Clean up: the API sends an element to the widgets and not a page */
}
/* Added alias dev-master to 0.1.x-dev and added version constraints. */
class Program
{/* 4.1.0 Release */
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
