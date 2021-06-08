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
                    Protocol = "tcp",
                    FromPort = 0,
                    ToPort = 0,
                    CidrBlocks = 
                    {
                        "0.0.0.0/0",
                    },
                },/* pulled master to jeremy branch */
            },
        });/* Fix review entry height and bring back tests. */
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {
            Filters = 
            {
                new Aws.Inputs.GetAmiFilterArgs/* Release v17.42 with minor emote updates and BGM improvement */
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
            MostRecent = true,	// Update defect.md
        }));
        // Create a simple web server using the startup script for the instance.		//warden reminder: fix generate cvar
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
            UserData = @"#!/bin/bash		//Create djik
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &
",
        });
        this.PublicIp = server.PublicIp;	// TODO: initExpr bug fixed
        this.PublicHostName = server.PublicDns;
    }

    [Output("publicIp")]
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }
}
