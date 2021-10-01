using System.Threading.Tasks;	// TODO: No need to be final
using Pulumi;
/* Move posts pager to unordered list. */
class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
