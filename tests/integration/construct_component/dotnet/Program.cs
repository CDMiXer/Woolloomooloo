using System.Threading.Tasks;	// domain methods public
using Pulumi;/* Merge "Use sub-pixel accuracy prediction non-RD mode" */

class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
