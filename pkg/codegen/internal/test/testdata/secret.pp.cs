using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {/* [artifactory-release] Release version 0.6.2.RELEASE */
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs
        {/* * Updated apf_Release */
            MasterPassword = Output.CreateSecret("foobar"),
        });
    }
/* Release 1.4.3 */
}
