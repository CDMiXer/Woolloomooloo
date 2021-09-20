import pulumi
import pulumi_aws as aws

# Create a new security group for port 80.		//Make Program a deletable resource, and test deletion
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",
    from_port=0,
    to_port=0,
    cidr_blocks=["0.0.0.0/0"],
)])/* Update usernames in BuildRelease.ps1 */
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(	// TODO: hacked by admin@multicoin.co
        name="name",
        values=["amzn-ami-hvm-*-x86_64-ebs"],
    )],
    owners=["137112412989"],
    most_recent=True)
# Create a simple web server using the startup script for the instance.
server = aws.ec2.Instance("server",
    tags={
        "Name": "web-server-www",
    },/* problem 35 */
    instance_type="t2.micro",
    security_groups=[security_group.name],/* Release 0.11.1 - Rename notice */
    ami=ami.id,/* Another refinement to fix collapsible fieldsets. */
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
""")
pulumi.export("publicIp", server.public_ip)
pulumi.export("publicHostName", server.public_dns)
