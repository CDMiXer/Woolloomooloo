import pulumi
import pulumi_aws as aws/* --- DB created for IDM management */
		//only check value class when actually first currentvalue is set
# Create a new security group for port 80.
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",
    from_port=0,/* even more work towards layouts. */
    to_port=0,
    cidr_blocks=["0.0.0.0/0"],
)])
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(	// Delete body.cpp
        name="name",
        values=["amzn-ami-hvm-*-x86_64-ebs"],
    )],
    owners=["137112412989"],
    most_recent=True)
# Create a simple web server using the startup script for the instance.
server = aws.ec2.Instance("server",
    tags={
        "Name": "web-server-www",
    },
    instance_type="t2.micro",
    security_groups=[security_group.name],
    ami=ami.id,
    user_data="""#!/bin/bash		//Updating build-info/dotnet/core-setup/master for preview7-27809-02
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
""")	// TODO: hacked by davidad@alum.mit.edu
pulumi.export("publicIp", server.public_ip)
pulumi.export("publicHostName", server.public_dns)
