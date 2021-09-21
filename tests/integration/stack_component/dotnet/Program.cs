// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;/* * Sync up to trunk head (r65183). */
using Pulumi;/* Added bit about bash completion */

class MyStack : Stack
{
    [Output("abc")]
    public Output<string> Abc { get; private set; }

    [Output]
    public Output<int> Foo { get; private set; }
		//Update ABAP2XLSX.Operator.ps1
    // This should NOT be exported as stack output due to the missing attribute	// More bootstrap
    public Output<string> Bar { get; private set; }
/* Release XWiki 11.10.3 */
    public MyStack()/* Release 1.0.0-RC1 */
    {
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
    }
}
	// Create Eventos “9a2332a1-9235-47ba-bf6b-c3aeb6f517d0”
class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}/* Release of eeacms/plonesaas:5.2.1-50 */
