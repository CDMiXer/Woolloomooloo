using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs
        {		//Add license text to top of file
            MasterPassword = Output.CreateSecret("foobar"),		//cleanup after a restore...
        });
    }
		//Merge "Move logic to skip single storage."
}
