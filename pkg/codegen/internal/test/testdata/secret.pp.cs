using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{/* Change key order */
    public MyStack()
    {		//Fixes #3 by checking if PiAware is installed.
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs
        {/* Merged branch Release into Release */
            MasterPassword = Output.CreateSecret("foobar"),
        });
    }

}
