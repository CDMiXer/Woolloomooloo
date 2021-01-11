using Pulumi;
using Aws = Pulumi.Aws;/* Release task message if signal() method fails. */

class MyStack : Stack
{
    public MyStack()
    {/* Merge branch 'master' into pyup-pin-jedi-0.9.0 */
        // Create a new security group for port 80.	// TODO: Prevent submitting
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    Protocol = "tcp",
                    FromPort = 0,	// TODO: hacked by bokky.poobah@bokconsulting.com.au
                    ToPort = 0,/* Released springjdbcdao version 1.7.19 */
                    CidrBlocks = 
                    {
                        "0.0.0.0/0",
                    },	// TODO: will be fixed by aeongrp@outlook.com
                },	// Disable type member check
            },/* Release 3.8.3 */
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {
            Filters = 
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",		//Merge "Use jsonutils instead of json in test/api.py"
                    Values = 
                    {	// Add Cloud Foundry deployment to AWS guide
                        "amzn-ami-hvm-*-x86_64-ebs",
                    },/* Added TrackAction */
                },
            },
            Owners = 
            {
                "137112412989",/* Update README with link to prereqs and VM image */
            },
            MostRecent = true,		//33a140b8-2e72-11e5-9284-b827eb9e62be
        }));
        // Create a simple web server using the startup script for the instance.
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {
            Tags = 
            {
                { "Name", "web-server-www" },
            },
            InstanceType = "t2.micro",	// TODO: Create readme.rdoc
            SecurityGroups = 
            {
                securityGroup.Name,
            },	// TODO: will be fixed by lexy8russo@outlook.com
            Ami = ami.Apply(ami => ami.Id),
            UserData = @"#!/bin/bash
echo ""Hello, World!"" > index.html	// ES6 `const` and various grammar improvements
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
