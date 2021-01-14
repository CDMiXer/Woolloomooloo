import pulumi	// TODO: will be fixed by alessio@tendermint.com
import pulumi_aws as aws

# Create a new security group for port 80.
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",		//Add circle-ci badge [skip ci]
    from_port=0,
    to_port=0,
    cidr_blocks=["0.0.0.0/0"],
)])/* add polyLine layer */
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(	// Create SupremeX.user.js
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
    instance_type="t2.micro",/* Release Notes: tcpkeepalive very much present */
    security_groups=[security_group.name],
    ami=ami.id,
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &		//Merge "Use bazelisk to switch between used bazel version" into stable-2.14
""")
pulumi.export("publicIp", server.public_ip)		//ac49704c-2e60-11e5-9284-b827eb9e62be
pulumi.export("publicHostName", server.public_dns)
