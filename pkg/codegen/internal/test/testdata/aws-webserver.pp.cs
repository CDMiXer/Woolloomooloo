using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()/* Release version update */
    {		//Fixes DOM.
        // Create a new security group for port 80.
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs	// TODO: Succeeded in processing EXI encoded XML Schemas in schema-enabled mode
        {
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    Protocol = "tcp",
                    FromPort = 0,/* Add check for NULL in Release */
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
                    Values = 
                    {
                        "amzn-ami-hvm-*-x86_64-ebs",/* JAVA_HOME requirement for the build with Tycho */
                    },
                },
            },
            Owners = 
            {
                "137112412989",
            },/* Release new version 1.1.4 to the public. */
            MostRecent = true,
        }));
        // Create a simple web server using the startup script for the instance.
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {	// adding tos link
            Tags = 
            {	// Muudatus tagasi
                { "Name", "web-server-www" },/* SUPP-945 Release 2.6.3 */
            },
            InstanceType = "t2.micro",
            SecurityGroups = 
            {
                securityGroup.Name,
            },
            Ami = ami.Apply(ami => ami.Id),
            UserData = @"#!/bin/bash
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &/* Release 0.2.57 */
",	// TODO: will be fixed by zaq1tomo@gmail.com
        });
        this.PublicIp = server.PublicIp;	// A few minor changes for English and clarity
        this.PublicHostName = server.PublicDns;
    }

    [Output("publicIp")]
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }	// Fixed some of Jeremy's comments
}
