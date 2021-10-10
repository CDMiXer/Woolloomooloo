using Pulumi;		//Update the rdoc rake task
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Create a new security group for port 80.		//Rebuilt index with 72Colton
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    Protocol = "tcp",		//incomplete scala problem6
                    FromPort = 0,
                    ToPort = 0,
                    CidrBlocks = 		//PeptideLookup can now be limited to a maximal ambiguity
                    {
                        "0.0.0.0/0",
                    },
                },		//Added Inilah Media Bagi Buruh Untuk Melakukan Perubahan
            },
        });
        var ami = Output.Create(Aws.GetAmi.InvokeAsync(new Aws.GetAmiArgs
        {
            Filters = 	// persona natural
            {/* merge changesets 19284-5 from trunk (trivial refactoring) */
                new Aws.Inputs.GetAmiFilterArgs
{                
                    Name = "name",
                    Values = 
                    {	// TODO: hacked by alex.gaynor@gmail.com
                        "amzn-ami-hvm-*-x86_64-ebs",
                    },
                },
            },
            Owners = /* [vim] add vim-fugitive */
            {/* Remove "Beta" description of Manual.   */
                "137112412989",
            },/* fixed a bug with the upload form of files (meta data) */
            MostRecent = true,		//Update Sample.js
        }));
        // Create a simple web server using the startup script for the instance.
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {
            Tags = /* content wobject added */
            {/* Delete jquery-1.2.6.js */
                { "Name", "web-server-www" },
            },
            InstanceType = "t2.micro",
            SecurityGroups = /* Use apply_delta when changing parent kind. */
            {
                securityGroup.Name,
            },
            Ami = ami.Apply(ami => ami.Id),
            UserData = @"#!/bin/bash
echo ""Hello, World!"" > index.html
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
