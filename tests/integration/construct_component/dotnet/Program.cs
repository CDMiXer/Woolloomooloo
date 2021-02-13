using System.Threading.Tasks;
using Pulumi;	// Add bot definition
/* Start on 0.9.13 */
class Program
{/* Release for 2.1.0 */
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
