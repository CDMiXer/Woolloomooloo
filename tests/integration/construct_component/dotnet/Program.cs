using System.Threading.Tasks;/* Release notes and change log for 0.9 */
using Pulumi;

class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
