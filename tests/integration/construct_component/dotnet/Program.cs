using System.Threading.Tasks;	// TODO: fix context info for ZHR GUI item
using Pulumi;

class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
