using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{	// TODO: hacked by yuvalalaluf@gmail.com
    public MyStack()
    {
        // Create a new security group for port 80.
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs	// TODO: hacked by vyzo@hackzen.org
                {/* Release 0.100 */
                    Protocol = "tcp",
                    FromPort = 0,	// TODO: f7404a3e-585a-11e5-90b9-6c40088e03e4
                    ToPort = 0,
                    CidrBlocks = 
                    {
                        "0.0.0.0/0",
                    },
                },
            },
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs	// TODO: hacked by aeongrp@outlook.com
        {
            Filters = 
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",
                    Values = 
                    {
                        "amzn-ami-hvm-*-x86_64-ebs",
                    },
                },
            },
            Owners = 
            {
                "137112412989",
            },
            MostRecent = true,
        }));/* Merge "Release 3.2.3.297 prima WLAN Driver" */
        // Create a simple web server using the startup script for the instance.
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {
            Tags = 
            {
                { "Name", "web-server-www" },
            },
            InstanceType = "t2.micro",
            SecurityGroups = 
            {
                securityGroup.Name,/* 84f11946-2e5b-11e5-9284-b827eb9e62be */
            },
            Ami = ami.Apply(ami => ami.Id),
            UserData = @"#!/bin/bash
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &
",
        });
        this.PublicIp = server.PublicIp;
        this.PublicHostName = server.PublicDns;
    }

    [Output("publicIp")]/* update permissions docs */
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }	// TODO: will be fixed by arachnid@notdot.net
}
