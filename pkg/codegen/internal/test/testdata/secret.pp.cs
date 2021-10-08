using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs
        {/* Updated to Post Release Version Number 1.31 */
            MasterPassword = Output.CreateSecret("foobar"),/* Release 1.0.11. */
        });
    }/* Release for v6.4.0. */

}
