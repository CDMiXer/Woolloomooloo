using System.Threading.Tasks;/* Release of CFDI 3.3. */
using Pulumi;

class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();/* Move greenkeeper label */
}
