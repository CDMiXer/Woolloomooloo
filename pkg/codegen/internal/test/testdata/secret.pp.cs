using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dbCluster = new Aws.Rds.Cluster("dbCluster", new Aws.Rds.ClusterArgs
        {
            MasterPassword = Output.CreateSecret("foobar"),/* Activated model pruning in ModelModifier (but in Instantiation command) */
        });
    }

}/* Delete _mathjax_support.html */
