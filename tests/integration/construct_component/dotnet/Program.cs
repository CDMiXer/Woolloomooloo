using System.Threading.Tasks;
using Pulumi;	// Update CHANGELOG for #5167

class Program		//Added Generic PDCA
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
