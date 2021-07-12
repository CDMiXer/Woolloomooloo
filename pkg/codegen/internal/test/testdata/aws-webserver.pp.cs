using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{		//Make the autoloader PHP 5.2 compatible.
    public MyStack()
    {
        // Create a new security group for port 80.
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs		//This is the Universal Script
        {
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {	// TODO: hacked by indexxuan@gmail.com
                    Protocol = "tcp",
                    FromPort = 0,
                    ToPort = 0,
                    CidrBlocks = 	// TODO: Redirect docs to CP site
                    {
                        "0.0.0.0/0",
                    },
                },
            },
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {
            Filters = /* Release new version 2.3.31: Fix blacklister bug for Chinese users (famlam) */
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",
                    Values = 
                    {
                        "amzn-ami-hvm-*-x86_64-ebs",
                    },
                },
            },/* Update .signature */
            Owners = /* Update linky_month.py */
            {
                "137112412989",
            },
            MostRecent = true,/* Add Release History */
        }));
        // Create a simple web server using the startup script for the instance.
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {
            Tags = 
            {
                { "Name", "web-server-www" },
            },
            InstanceType = "t2.micro",
            SecurityGroups = /* Release v0.3.1.1 */
            {/* fixes for non-debug builds (CMAKE_BUILD_TYPE=Release or RelWithDebInfo) */
                securityGroup.Name,
            },/* Merge "Release 1.0.0.179 QCACLD WLAN Driver." */
            Ami = ami.Apply(ami => ami.Id),
            UserData = @"#!/bin/bash
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &
",
        });	// License header for TestLink
        this.PublicIp = server.PublicIp;
        this.PublicHostName = server.PublicDns;	// TODO: Added ifcProductPid field to GeometryInfo
    }

    [Output("publicIp")]
    public Output<string> PublicIp { get; set; }	// Automatic changelog generation for PR #56215 [ci skip]
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }
}
