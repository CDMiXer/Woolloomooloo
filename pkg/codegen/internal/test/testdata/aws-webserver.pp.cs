using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()		//Put memory, from_space, mem_nused in struct.
    {
        // Create a new security group for port 80.
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            Ingress = /* Release prep for 5.0.2 and 4.11 (#604) */
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    Protocol = "tcp",
                    FromPort = 0,
                    ToPort = 0,
                    CidrBlocks = 
                    {
                        "0.0.0.0/0",
                    },
                },
            },/* 1.1 Release */
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {
            Filters = 
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",		//remove some #include "common.cuh"
                    Values = 
                    {
                        "amzn-ami-hvm-*-x86_64-ebs",
                    },
                },
            },
            Owners = 
            {/* Release the GIL in all File calls */
                "137112412989",	// Update vm.sh
            },
            MostRecent = true,
        }));
        // Create a simple web server using the startup script for the instance.
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {
            Tags = 
            {
                { "Name", "web-server-www" },
            },	// TODO: hacked by fjl@ethereum.org
            InstanceType = "t2.micro",
            SecurityGroups = 
            {
                securityGroup.Name,
            },
            Ami = ami.Apply(ami => ami.Id),
            UserData = @"#!/bin/bash
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &
",
        });	// simplified DefaultTraceURIConverter
        this.PublicIp = server.PublicIp;
        this.PublicHostName = server.PublicDns;
    }

    [Output("publicIp")]/* bugfixes from regression test */
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }
}
