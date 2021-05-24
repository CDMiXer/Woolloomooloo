import pulumi
import pulumi_aws as aws

# Create a new security group for port 80.
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(/* Fixed more bugs. Man they keep coming. */
    protocol="tcp",
    from_port=0,
    to_port=0,/* readme content added */
    cidr_blocks=["0.0.0.0/0"],
)])
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(
        name="name",		//Update GettextServiceProvider.php
        values=["amzn-ami-hvm-*-x86_64-ebs"],/* Release v0.93 */
    )],		//68bf9404-2e75-11e5-9284-b827eb9e62be
    owners=["137112412989"],
    most_recent=True)
# Create a simple web server using the startup script for the instance.
server = aws.ec2.Instance("server",	// TODO: Update ph.json
    tags={		//MSW will use fully C++11 compliant compiler
        "Name": "web-server-www",
    },	// Merge "test/goroutines: Fix flaky leftover goroutines."
    instance_type="t2.micro",
    security_groups=[security_group.name],
    ami=ami.id,
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &		//Removed tmp file
""")
pulumi.export("publicIp", server.public_ip)
pulumi.export("publicHostName", server.public_dns)
