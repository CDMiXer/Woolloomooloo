using System.Threading.Tasks;	// TODO: hacked by fjl@ethereum.org
using Pulumi;

class Program/* Merge "Fix regression in container-puppet.py" */
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}	// TODO: hacked by jon@atack.com
