using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()		//Merge branch 'develop' into bug/remove-view-wallet
    {
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs	// Refactor to eliminate logger factory calls that are direct
        {
            MasterPassword = Output.CreateSecret("foobar"),
        });
    }

}	// Merge "Enable preloading dashboard queries on root url"
