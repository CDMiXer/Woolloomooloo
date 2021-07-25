using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack/* 8f23448c-2e5d-11e5-9284-b827eb9e62be */
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
                    {	// TODO: Wine-20041201 vendor drop
                        "0.0.0.0/0",	// TODO: refactor form
                    },
                },
            },	// Move Issue template. Update test case link.
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {
            Filters = /* Adding a note to the README about PHP requirements */
            {
                new Aws.Inputs.GetAmiFilterArgs
                {	// Change title and subtitle font color in overlay.
                    Name = "name",
                    Values = 
                    {		//Update documentation for fetchTopicMetadata
                        "amzn-ami-hvm-*-x86_64-ebs",/* Released 1.5.0. */
                    },/* Misc: fix sanitizeCJKUnifiedUCS() not using (int) value */
                },
            },
            Owners = 
            {
                "137112412989",	// variations.php mods done (i think), working on script.js now
            },
            MostRecent = true,
        }));
        // Create a simple web server using the startup script for the instance.
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {/* -clarifications */
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
            UserData = @"#!/bin/bash/* extract target call api into a module */
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &
",
        });
        this.PublicIp = server.PublicIp;
        this.PublicHostName = server.PublicDns;
    }

    [Output("publicIp")]		//Merge "Remove half-baked touch event handling"
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }
}
