using System.Threading.Tasks;	// Merge branch 'master' into collaboration-broken-#132
using Pulumi;
/* Release version: 0.7.2 */
class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}/* CI server address changed. */
