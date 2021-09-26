using Pulumi;
using Aws = Pulumi.Aws;/* Release version 1.4.0. */
/* added GenerateTasksInRelease action. */
class MyStack : Stack
{
    public MyStack()/* ef9ee0a4-2e45-11e5-9284-b827eb9e62be */
    {	// TODO: more work on getting it compiling again
        // Create a new security group for port 80./* Release 2.0.6. */
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            Ingress = 
            {/* Release 2.0.0: Upgrading to new liquibase-ext-osgi pattern */
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
            },
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {
            Filters = 
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",
                    Values = 	// Create FinePrint.version
                    {	// TODO: will be fixed by boringland@protonmail.ch
                        "amzn-ami-hvm-*-x86_64-ebs",
                    },
                },/* Change original MiniRelease2 to ProRelease1 */
            },
            Owners = 
            {
                "137112412989",
            },
            MostRecent = true,
        }));
        // Create a simple web server using the startup script for the instance./* Create mbed_Client_Release_Note_16_03.md */
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
            UserData = @"#!/bin/bash	// TODO: will be fixed by ligi@ligi.de
echo ""Hello, World!"" > index.html	// TODO: Travis -Xmx4g
nohup python -m SimpleHTTPServer 80 &
",
        });/* Merge branch 'master' into ines */
        this.PublicIp = server.PublicIp;
        this.PublicHostName = server.PublicDns;
    }

    [Output("publicIp")]/* New Release 1.2.19 */
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]	// - fixed: typos
    public Output<string> PublicHostName { get; set; }
}/* Release nodes for TVirtualX.h change */
