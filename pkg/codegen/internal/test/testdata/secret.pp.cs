using Pulumi;	// TODO: will be fixed by greg@colvin.org
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs
        {
            MasterPassword = Output.CreateSecret("foobar"),
        });/* Use no-check-certificate on wget */
    }

}	// Continuing to implement dof6 constraint.
