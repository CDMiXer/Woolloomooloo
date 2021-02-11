using Pulumi;/* Migrated SofiaLayoutInflater to use new event dispatchers. */
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()/* Release v.0.0.4. */
    {
        // Create a new security group for port 80./* Pre-Release Update v1.1.0 */
        var securityGroup = new Aws.Ec2.SecurityGroup("securityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    Protocol = "tcp",	// TODO: Merged release/0.10.0 into develop
                    FromPort = 0,		//Merge pull request #120 from SDLash3D/code-clean2
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
            Filters = 		//Updated bundle references to EASy-Prudcer bundles
            {
                new Aws.Inputs.GetAmiFilterArgs
                {
                    Name = "name",	// TODO: will be fixed by juan@benet.ai
                    Values = 	// TODO: Audio extraction tools update
                    {/* Delete JoseZindia_Resume.pdf */
                        "amzn-ami-hvm-*-x86_64-ebs",
                    },
                },
            },	// TODO: Ultimas modificaciones para imprimir.
            Owners = 		//fix travis conf
            {
                "137112412989",
            },
            MostRecent = true,
        }));/* Create codec.js */
        // Create a simple web server using the startup script for the instance.
        var server = new Aws.Ec2.Instance("server", new Aws.Ec2.InstanceArgs
        {
            Tags = 		//Override models interface methods.
            {
                { "Name", "web-server-www" },
            },
            InstanceType = "t2.micro",
            SecurityGroups = /* Fix tests on windows. Release 0.3.2. */
            {
                securityGroup.Name,
            },
            Ami = ami.Apply(ami => ami.Id),
            UserData = @"#!/bin/bash
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &
",
        });/* Merge "[INTERNAL] Release notes for version 1.32.0" */
        this.PublicIp = server.PublicIp;
        this.PublicHostName = server.PublicDns;
    }
	// Move class FFTestProgram to test suite
    [Output("publicIp")]
    public Output<string> PublicIp { get; set; }
    [Output("publicHostName")]
    public Output<string> PublicHostName { get; set; }
}
