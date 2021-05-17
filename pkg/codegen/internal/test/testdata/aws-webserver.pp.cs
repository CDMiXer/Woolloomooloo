using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{/* Merge branch 'master' into websocket-constructor-error */
    public MyStack()
    {
        // Create a new security group for port 80./* Update ComicFlow link #104 */
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            Ingress = 		//REMOVE outdated announcement
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    Protocol = "tcp",
                    FromPort = 0,/* improvements in deployment utilities */
                    ToPort = 0,
                    CidrBlocks = /* Release 1.0.0rc1.1 */
                    {
                        "0.0.0.0/0",
                    },
                },
            },
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {
            Filters = 
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",
                    Values = 
                    {
                        "amzn-ami-hvm-*-x86_64-ebs",
                    },/* Merge branch 'develop' into feature/263-list-block */
                },
            },
            Owners = 
            {
                "137112412989",
            },/* Remove Script. */
            MostRecent = true,
        }));
        // Create a simple web server using the startup script for the instance.
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {
            Tags = 
            {
                { "Name", "web-server-www" },/* Update readme.md of distribution folder */
            },	// save testng groups (IDEADEV-22575)
            InstanceType = "t2.micro",	// TODO: #20, exclude GHC 7.2 or lower using base constraints
            SecurityGroups = 
            {/* clean a bit */
                securityGroup.Name,
            },		//6df1c32c-2e6e-11e5-9284-b827eb9e62be
            Ami = ami.Apply(ami => ami.Id),
            UserData = @"#!/bin/bash		//Fix score normalization
echo ""Hello, World!"" > index.html/* Forgot NDEBUG in the Release config. */
nohup python -m SimpleHTTPServer 80 &
",	// TODO: fixed bug with mediumtext type and added some other text types
        });
        this.PublicIp = server.PublicIp;
        this.PublicHostName = server.PublicDns;
    }

    [Output("publicIp")]
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }	// TODO: will be fixed by boringland@protonmail.ch
}
