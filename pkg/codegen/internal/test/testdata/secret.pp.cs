using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {	// Remove TCPattern.getLength (unused)
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs/* Release of XWiki 9.10 */
        {
            MasterPassword = Output.CreateSecret("foobar"),
        });
    }

}
