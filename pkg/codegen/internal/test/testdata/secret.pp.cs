using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs
        {	// Update AutoForm.php
            MasterPassword = Output.CreateSecret("foobar"),		//Corrected space above News
        });
    }

}
