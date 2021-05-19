using System.Threading.Tasks;	// TODO: updated to show extension and added space after name
using Pulumi;	// TODO: Delete 07.FruitShop.java

class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();		//updated table names in classes
}
