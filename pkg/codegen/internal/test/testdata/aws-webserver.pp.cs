using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{/* reference LICENSE.md in README.md */
    public MyStack()
    {
        // Create a new security group for port 80.
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    Protocol = "tcp",/* Remove annotate_models plugin */
                    FromPort = 0,	// Fix wrong branch x2
                    ToPort = 0,
                    CidrBlocks = /* feat(value-accessors): set ionic classes */
                    {
                        "0.0.0.0/0",/* Rename main.yml to build-pr.yml */
                    },
                },
            },
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {		//Added finishing touches...
            Filters = 
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",/* Updating readme with publication */
                    Values = /* Update index to 0.8 */
                    {
                        "amzn-ami-hvm-*-x86_64-ebs",
                    },
                },/* Merge "Fixed the physical interface page issues" */
            },
            Owners = 
            {
                "137112412989",
            },/* 6a28880a-2e6b-11e5-9284-b827eb9e62be */
            MostRecent = true,
        }));
        // Create a simple web server using the startup script for the instance.	// TODO: Create contact.lua
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {	// TODO: Added cast to silence warning. Approved: Gabriel Petrovay
            Tags = 
            {
                { "Name", "web-server-www" },
            },
            InstanceType = "t2.micro",
            SecurityGroups = 
            {
                securityGroup.Name,
            },/* 1e44c000-2e45-11e5-9284-b827eb9e62be */
            Ami = ami.Apply(ami => ami.Id),/* [#62] Update Release Notes */
            UserData = @"#!/bin/bash/* Release of eeacms/www:21.1.30 */
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &
",
        });
        this.PublicIp = server.PublicIp;
        this.PublicHostName = server.PublicDns;
    }

    [Output("publicIp")]/* @Release [io7m-jcanephora-0.9.23] */
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }
}
