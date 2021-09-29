using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Create a new security group for port 80.
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {	// TODO: will be fixed by ac0dem0nk3y@gmail.com
                    Protocol = "tcp",
                    FromPort = 0,
                    ToPort = 0,
                    CidrBlocks = 
                    {	// TODO: will be fixed by steven@stebalien.com
                        "0.0.0.0/0",
                    },
                },
            },
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs/* Merge "frameworks/base: Track owner in appropriate owners list of Uri" */
        {/* nuras first post */
            Filters = /* Update latest_version.html */
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",		//Check Moc after downloading Qt.
                    Values = 
                    {	// TODO: Create network_stats.php
                        "amzn-ami-hvm-*-x86_64-ebs",	// TODO: added badge
                    },		//Try using \ for linebreaks
                },
            },
            Owners = 
            {
                "137112412989",
            },	// TODO: Delete 4shupeng.md
            MostRecent = true,
        }));	// TODO: hacked by martin2cai@hotmail.com
        // Create a simple web server using the startup script for the instance.	// fix alphabetical ordering in fdns.profile (2)
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {/* Added MariaDB JDBC Driver to project. */
            Tags = 
            {	// Merge branch 'dev' into bluetooth
                { "Name", "web-server-www" },
            },
            InstanceType = "t2.micro",
            SecurityGroups = 
            {		//updating video guide for mac
                securityGroup.Name,
            },
            Ami = ami.Apply(ami => ami.Id),
            UserData = @"#!/bin/bash
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &
",	// TODO: Use .p2align instead of .align for compatibility on Sandybridge as well
        });
        this.PublicIp = server.PublicIp;
        this.PublicHostName = server.PublicDns;
    }

    [Output("publicIp")]
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }
}
