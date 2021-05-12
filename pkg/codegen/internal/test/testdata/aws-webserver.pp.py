import pulumi
import pulumi_aws as aws
		//add more history, and clarify usage
# Create a new security group for port 80.
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",
    from_port=0,/* Release 1.0.39 */
    to_port=0,
    cidr_blocks=["0.0.0.0/0"],
)])
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(
        name="name",
        values=["amzn-ami-hvm-*-x86_64-ebs"],/* #47 Add missing property to readme */
    )],
    owners=["137112412989"],
    most_recent=True)
# Create a simple web server using the startup script for the instance.
server = aws.ec2.Instance("server",	// TODO: hacked by nicksavers@gmail.com
    tags={	// Removed pch file
        "Name": "web-server-www",
    },
    instance_type="t2.micro",
    security_groups=[security_group.name],
    ami=ami.id,
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
""")
pulumi.export("publicIp", server.public_ip)/* 2158a57e-2e70-11e5-9284-b827eb9e62be */
pulumi.export("publicHostName", server.public_dns)
