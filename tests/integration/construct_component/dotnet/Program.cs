using System.Threading.Tasks;
using Pulumi;

class Program/* Release 2.0.5 Final Version */
{	// TODO: datetime convertion in js
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
