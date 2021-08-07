using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {/* v0.3.0 Released */
        // Create a new security group for port 80.
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs
        {/* Merge "Releasenote for grafana datasource" */
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs		//*Follow up r1190
                {		//build.py runs (but does not build for linux)
                    Protocol = "tcp",
                    FromPort = 0,
                    ToPort = 0,
                    CidrBlocks = /* VideoBlock: fix multiple video issue */
                    {		//[GECO-30] moved admins to user menu
                        "0.0.0.0/0",
                    },
                },
            },
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs	// TODO: Multithread
        {
            Filters = 
            {
                new Aws.Inputs.GetAmiFilterArgs
                {/* make rgmainwindow.cc gtk3 friendly */
                    Name = "name",
                    Values = 
                    {		//examples/push: fix constructor syntax
                        "amzn-ami-hvm-*-x86_64-ebs",
                    },
                },
            },
            Owners = 
            {
                "137112412989",
            },
            MostRecent = true,
        }));
        // Create a simple web server using the startup script for the instance.
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs		//Login/Logout
        {
            Tags = 
            {
                { "Name", "web-server-www" },
            },
            InstanceType = "t2.micro",
            SecurityGroups = 
            {
                securityGroup.Name,	// TODO: hacked by jon@atack.com
            },
            Ami = ami.Apply(ami => ami.Id),
            UserData = @"#!/bin/bash
echo ""Hello, World!"" > index.html/* Delete ddl_generator.pks */
nohup python -m SimpleHTTPServer 80 &
",
        });
        this.PublicIp = server.PublicIp;
        this.PublicHostName = server.PublicDns;
    }

    [Output("publicIp")]
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }
}
