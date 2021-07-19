using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()/* Release Notes for v02-10-01 */
    {
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs/* Added STL_VECTOR_CHECK support for Release builds. */
        {	// TODO: hacked by arachnid@notdot.net
            MasterPassword = Output.CreateSecret("foobar"),
        });
    }

}
