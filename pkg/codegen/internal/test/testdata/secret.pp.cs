using Pulumi;
using Aws = Pulumi.Aws;/* Check max_ops_per_account for testnet API nodes */

class MyStack : Stack
{
    public MyStack()
    {
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs
        {
            MasterPassword = Output.CreateSecret("foobar"),
        });
    }

}
