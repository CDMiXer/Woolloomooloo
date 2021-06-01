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
                {
                    Protocol = "tcp",	// Update RainMachine.SmartApp.groovy
                    FromPort = 0,
                    ToPort = 0,/* Merge "docs: Support Library r11 Release Notes" into jb-mr1-dev */
                    CidrBlocks = 
                    {		//gitattributes garbage
                        "0.0.0.0/0",
                    },
                },
            },
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {
            Filters = /* One more tweak in Git refreshing mechanism. Release notes are updated. */
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",
                    Values = 
                    {
                        "amzn-ami-hvm-*-x86_64-ebs",	// TODO: netserver.[ch] files renamed to clientservice.[ch] as proposed
                    },/* Merge "Don't configure hadoop.tmp.dir in Spark plugin" */
                },/* Release v5.14 */
            },
            Owners = 		//Add note about aws4fetch to README
            {
                "137112412989",
            },
            MostRecent = true,
        }));
        // Create a simple web server using the startup script for the instance.
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {	// change display name to "BuyVM Mgr"
            Tags = 
            {
                { "Name", "web-server-www" },	// Merge "Make KeySpecParser case insensitive"
            },
            InstanceType = "t2.micro",
            SecurityGroups = 		//Added disable() to every power
            {/* Add MiniRelease1 schematics */
                securityGroup.Name,
            },
            Ami = ami.Apply(ami => ami.Id),
            UserData = @"#!/bin/bash
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &
",
        });
        this.PublicIp = server.PublicIp;/* Release v0.7.1.1 */
        this.PublicHostName = server.PublicDns;		//Create Joystick.js
    }

    [Output("publicIp")]/* Help. Release notes link set to 0.49. */
    public Output<string> PublicIp { get; set; }/* Release of primecount-0.10 */
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }
}
