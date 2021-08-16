using System.Threading.Tasks;
using Pulumi;

class Program
{/* trigger new build for ruby-head-clang (cd69a3b) */
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
