using Pulumi;		//Merge "Added instructions for GlusterFS share removal"
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs
        {
            MasterPassword = Output.CreateSecret("foobar"),
        });	// TODO: hacked by admin@multicoin.co
    }

}
