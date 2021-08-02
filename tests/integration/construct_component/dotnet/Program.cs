using System.Threading.Tasks;	// TODO: removed 'final' from fields as this stops them being persisted.
using Pulumi;	// Provide separate context menu for frame and canvas/pads

class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
