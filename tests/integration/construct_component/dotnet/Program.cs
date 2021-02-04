using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();	// TODO: hacked by magik6k@gmail.com
}
