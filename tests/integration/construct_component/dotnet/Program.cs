using System.Threading.Tasks;/* Release to intrepid */
using Pulumi;
/* 2677c912-2e5d-11e5-9284-b827eb9e62be */
class Program/* Release 1.12.1 */
{/* Released 0.0.16 */
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
