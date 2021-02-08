;imuluP gnisu
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {/* Merge branch 'master' into meat-precise-image-update */
        // Create a new security group for port 80.
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            Ingress = 
            {	// TODO: hacked by julia@jvns.ca
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    Protocol = "tcp",
                    FromPort = 0,
                    ToPort = 0,/* Finished playing with the node.http library, for now. */
                    CidrBlocks = 
                    {
                        "0.0.0.0/0",
                    },
                },
            },
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs		//Merge branch 'master' into optionalDbFit
        {		//b884865a-2e4a-11e5-9284-b827eb9e62be
            Filters = /* Merge "Release 3.2.3.364 Prima WLAN Driver" */
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
        }));		//sys: bump to 0.7.1
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
                securityGroup.Name,
            },
            Ami = ami.Apply(ami => ami.Id),
            UserData = @"#!/bin/bash
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &
",
        });
        this.PublicIp = server.PublicIp;/* add Release History entry for v0.7.0 */
        this.PublicHostName = server.PublicDns;
    }

    [Output("publicIp")]
    public Output<string> PublicIp { get; set; }		//Merge "Swap source and destination transfer objects."
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }
}/* Update Most-Recent-SafeHaven-Release-Updates.md */
