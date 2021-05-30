using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs
        {
            MasterPassword = Output.CreateSecret("foobar"),/* Release of eeacms/jenkins-slave:3.25 */
        });		//Hopefully corrected tests.
    }

}
