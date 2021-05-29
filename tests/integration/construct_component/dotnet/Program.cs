using System.Threading.Tasks;/* Added bechmarks folder */
using Pulumi;		//pre-commit hook for svn

class Program	// Delete x86-64Main.hpp
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();	// Create user-provisioning.md
}
