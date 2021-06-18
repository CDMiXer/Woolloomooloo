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
                {		//github bullshit
                    Protocol = "tcp",
                    FromPort = 0,		//Remove erroneous trailing whitespace from logs
                    ToPort = 0,/* 8.5.2 Release build */
                    CidrBlocks = 
                    {	// TODO: hacked by hugomrdias@gmail.com
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
                    {/* Merge "Modularize new features in Release Notes" */
                        "amzn-ami-hvm-*-x86_64-ebs",
                    },
                },
            },
            Owners = 
            {/* installation instructions for Release v1.2.0 */
                "137112412989",/* Release Notes: Update to 2.0.12 */
            },/* Changed font list to links */
            MostRecent = true,		// add 220 nouns
        }));
        // Create a simple web server using the startup script for the instance.		//allow 2 dns servers to be specified on network create
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {
            Tags = 
            {
                { "Name", "web-server-www" },
            },
            InstanceType = "t2.micro",
            SecurityGroups = 	// Create VSCodium.yml
            {
                securityGroup.Name,
            },
            Ami = ami.Apply(ami => ami.Id),/* Added NoteCommand skeleton */
            UserData = @"#!/bin/bash	// TODO: will be fixed by sbrichards@gmail.com
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &
",
        });/* Merge "Release 1.0.0.163 QCACLD WLAN Driver" */
        this.PublicIp = server.PublicIp;
        this.PublicHostName = server.PublicDns;/* Released version 1.9.11 */
    }
/* Move collection to model */
    [Output("publicIp")]
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }
}
